package observation

import (
	"context"

	"github.com/shinto-dev/url-shortener/foundation/observation/apm"
	"github.com/shinto-dev/url-shortener/foundation/observation/logging"
	"github.com/shinto-dev/url-shortener/foundation/observation/trace"
)

type FieldFn func(ctx context.Context)

func Add(ctx context.Context, fieldFns ...FieldFn) {
	for _, fieldFn := range fieldFns {
		fieldFn(ctx)
	}
}

type Config struct {
	Context        string
	TraceID        string
	SupportLogging bool
	SupportAPM     bool
}

func WithObservation(ctx context.Context, config Config) context.Context {
	if config.SupportLogging {
		ctx = logging.WithLogger(ctx)
		Add(ctx, logging.LField("context", config.Context))
	}

	if config.SupportAPM {
		ctx = apm.WithAPM(ctx, config.Context)
	}

	if config.TraceID != "" {
		ctx = trace.WithTraceID(ctx, config.Context)
	}

	return ctx
}
