package config

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ksysoev/deriv-api-bff/pkg/core/handlerfactory"
	"go.etcd.io/etcd/clientv3"
)

const defaultTimeoutSeconds = 5

type EtcdConfig struct {
	KeyPrefix string `mapstructure:"key_prefix"`
	Servers   string `mapstructure:"servers"`
}

type EtcdSource struct {
	prefix string
	mu     sync.RWMutex
	cli    *clientv3.Client
}

func NewEtcdSource(cfg EtcdConfig) (*EtcdSource, error) {
	serves := strings.Split(cfg.Servers, ",")

	if len(serves) == 0 {
		return nil, fmt.Errorf("no etcd servers provided")
	}

	if cfg.KeyPrefix == "" {
		return nil, fmt.Errorf("no etcd key prefix provided")
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   serves,
		DialTimeout: defaultTimeoutSeconds * time.Second,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %w", err)
	}

	return &EtcdSource{
		prefix: cfg.KeyPrefix,
		cli:    cli,
	}, nil
}

func (es *EtcdSource) LoadConfig(ctx context.Context) ([]handlerfactory.Config, error) {
	data, err := es.cli.Get(ctx, es.prefix, clientv3.WithPrefix())

	if err != nil {
		return nil, fmt.Errorf("failed to get config from etcd: %w", err)
	}

	if data.Count == 0 {
		return nil, fmt.Errorf("no config found")
	}

	cfg := make([]handlerfactory.Config, 0, data.Count)

	for _, kv := range data.Kvs {
		var c handlerfactory.Config
		err := json.Unmarshal(kv.Value, &c)

		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal config: %w", err)
		}

		cfg = append(cfg, c)
	}

	return cfg, nil
}

func (es *EtcdSource) PutConfig(ctx context.Context, cfg []handlerfactory.Config) error {
	//TODO: add logic for removing keys that are not in the new config

	for _, c := range cfg {
		data, err := json.Marshal(c)

		if err != nil {
			return fmt.Errorf("failed to marshal config: %w", err)
		}

		_, err = es.cli.Put(ctx, es.prefix+c.Method, string(data))

		if err != nil {
			return fmt.Errorf("failed to put config: %w", err)
		}
	}

	return nil
}
