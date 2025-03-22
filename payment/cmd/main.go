package main

import (
	"github.com/abdelrahmanShawki/eSHOP/order/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
)

const (
	service     = "payment"
	environment = "dev"
	id          = 2
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func init() {
	log.SetFormatter(customLogger{
		formatter: log.JSONFormatter{FieldMap: log.FieldMap{
			"msg": "message",
		}},
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

type customLogger struct {
	formatter log.JSONFormatter
}

func (l customLogger) Format(entry *log.Entry) ([]byte, error) {
	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	//Below injection is Just to understand what Context has
	entry.Data["Context"] = span.SpanContext()
	return l.formatter.Format(entry)
}

func main() {
	tp, err := tracerProvider("http://jaeger-otel.jaeger.svc.cluster.local:14278/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
