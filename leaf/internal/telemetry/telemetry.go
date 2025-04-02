package telemetry

import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
    "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
    "go.opentelemetry.io/otel/metric/global"
    "go.opentelemetry.io/otel/propagation"
    "go.opentelemetry.io/otel/sdk/metric"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
    "log"
)

func InitTracer(serviceName string) (*trace.TracerProvider, error) {
    // Eksporter OTLP (gRPC) dla tracingu
    traceExporter, err := otlptrace.New(
        context.Background(),
        otlptracegrpc.NewClient(
            otlptracegrpc.WithEndpoint("localhost:4317"), // Adres OTLP Collector
            otlptracegrpc.WithInsecure(), // Użyj WithInsecure tylko dla testów lokalnych
        ),
    )
    if err != nil {
        return nil, err
    }

    // Provider tracingu
    tp := trace.NewTracerProvider(
        trace.WithBatcher(traceExporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
    )

    // Ustawienie globalnego providera
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.TraceContext{})

    return tp, nil
}

func InitMeter(serviceName string) (*metric.MeterProvider, error) {
    // Eksporter OTLP (gRPC) dla metryk
    metricExporter, err := otlpmetric.New(
        context.Background(),
        otlpmetricgrpc.NewClient(
            otlpmetricgrpc.WithEndpoint("localhost:4317"), // Adres OTLP Collector
            otlpmetricgrpc.WithInsecure(), // Użyj WithInsecure tylko dla testów lokalnych
        ),
    )
    if err != nil {
        return nil, err
    }

    // Provider metryk
    mp := metric.NewMeterProvider(
        metric.WithReader(metric.NewPeriodicReader(metricExporter)),
        metric.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
    )

    // Ustawienie globalnego providera
    global.SetMeterProvider(mp)

    return mp, nil
}