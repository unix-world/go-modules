package tollbooth_chi

import (
	"net/http"

	"github.com/unix-world/go-modules/didip/tollbooth"
	"github.com/unix-world/go-modules/didip/tollbooth/config"
)

func LimitHandler(limiter *config.Limiter) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		wrapper := &limiterWrapper{
			limiter: limiter,
		}

		wrapper.handler = handler
		return wrapper
	}
}

type limiterWrapper struct {
	limiter *config.Limiter
	handler http.Handler
}

func (l *limiterWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		http.Error(w, "Context was canceled", http.StatusServiceUnavailable)
		return

	default:
		httpError := tollbooth.LimitByRequest(l.limiter, r)
		if httpError != nil {
			w.Header().Add("Content-Type", l.limiter.MessageContentType)
			w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}

		l.handler.ServeHTTP(w, r)
	}
}
