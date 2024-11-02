package middleware

import (
	"time"

	"github.com/ksysoev/wasabi"
	"github.com/ksysoev/wasabi/dispatch"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func NewRedMetricsMiddleware(scope string) func(next wasabi.RequestHandler) wasabi.RequestHandler {
	meter := otel.GetMeterProvider().Meter(scope)
	timing, err := meter.Float64Histogram(
		"request_duration",
		metric.WithDescription("Request duration"),
		metric.WithUnit("milliseconds"),
	)

	if err != nil {
		panic("failed to initialize metric" + err.Error())
	}

	return func(next wasabi.RequestHandler) wasabi.RequestHandler {
		return dispatch.RequestHandlerFunc(func(conn wasabi.Connection, req wasabi.Request) error {
			start := time.Now()

			err := next.Handle(conn, req)

			elapsed := time.Since(start)

			isError := "f"
			if err != nil {
				isError = "t"
			}

			timing.Record(
				req.Context(),
				float64(elapsed.Milliseconds()),
				metric.WithAttributes(
					attribute.String("routing_key", req.RoutingKey()),
					attribute.String("error", isError),
				),
			)

			return err
		})
	}
}
