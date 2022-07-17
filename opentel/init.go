package opentel

import (
	"context"

	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var GinTracer = otel.Tracer("gin-server")

var TracerProvider *sdktrace.TracerProvider

func InitTracer() (disconnect func() error, err error) {
	exporter, err := stdout.New()
	if err != nil {
		return nil, err
	}

	TracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(TracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return func() error {
		return TracerProvider.Shutdown(context.Background())
	}, nil
}
