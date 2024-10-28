package config

// defines the source of the config and/or config changes“
type Source interface {
	Init() error

	GetConfigurations() (*Config, error)

	WatchConfig(string) error

	GetPriority() Priority

	Name() string

	Close()
}

type Priority int8

const (
	P0 Priority = iota
	P1
	P2
)