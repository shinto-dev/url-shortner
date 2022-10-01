package web

import (
	"net/http"
	"url-shortener/platform/observation"
	"url-shortener/platform/observation/apm"
	"url-shortener/platform/observation/logging"
	"url-shortener/platform/observation/trace"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func HandleRequest(apiContext string, fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := logging.WithLogger(r.Context())
		ctx = trace.WithTraceID(ctx, TraceID(r))
		ctx = apm.WithAPM(ctx, apiContext)
		observation.Add(ctx, logging.LField("context", apiContext))

		segment := apm.StartSegment(ctx, "handler")
		defer segment.ObserveDuration()

		if err := fn(w, r.WithContext(ctx)); err != nil {
			HandleError(ctx, w, err)
		}

		logging.FromContext(ctx).Info("Request completed")
	}
}
