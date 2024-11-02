package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

type OtelConfig struct {
	Prometheus *PrometheusConfig `mapstructure:"prometheus"`
}

type PrometheusConfig struct {
	Listen string `mapstructure:"listen"`
	Path   string `mapstructure:"path"`
}

func initMetricProvider(ctx context.Context, cfg *OtelConfig) error {
	if cfg.Prometheus != nil {
		if err := initPrometheus(ctx, cfg.Prometheus); err != nil {
			return fmt.Errorf("failed to initialize Prometheus: %w", err)
		}
	}

	return nil
}

func initPrometheus(ctx context.Context, cfg *PrometheusConfig) error {
	if cfg.Listen == "" {
		return fmt.Errorf("prometheus listen address is required")
	}
	if cfg.Path == "" {
		return fmt.Errorf("prometheus metrics path is required")
	}

	metricExporter, err := prometheus.New()
	if err != nil {
		return fmt.Errorf("failed to create Prometheus exporter: %w", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metricExporter),
	)

	otel.SetMeterProvider(meterProvider)

	go func() {
		if err := servePrometheus(ctx, cfg); err != nil {
			slog.Error("failed to serve Prometheus", slog.Any("error", err))
		}
	}()

	return nil
}

func servePrometheus(ctx context.Context, cfg *PrometheusConfig) error {
	mux := http.NewServeMux()
	mux.Handle(cfg.Path, promhttp.Handler())

	httpSrv := &http.Server{
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	slog.Info("serving metrics", slog.Any("listen", cfg.Listen), slog.Any("path", cfg.Path))

	lis, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()

		if err := httpSrv.Close(); err != nil {
			slog.Error("failed to close metric server", slog.Any("error", err))
		}
	}()

	return httpSrv.Serve(lis)
}
