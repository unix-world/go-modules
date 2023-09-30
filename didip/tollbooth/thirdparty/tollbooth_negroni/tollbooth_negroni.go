package tollbooth_negroni

import (
	"net/http"

	"github.com/syreclabs/tollbooth"
	"github.com/syreclabs/tollbooth/config"
	"github.com/urfave/negroni"
)

func LimitHandler(limiter *config.Limiter) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		httpError := tollbooth.LimitByRequest(limiter, r)
		if httpError != nil {
			w.Header().Add("Content-Type", limiter.MessageContentType)
			/* RHMOD Fix for error "http: multiple response.WriteHeader calls"
			   Reverse the sequence of the functions calls w.WriteHeader() and w.Write()
			*/
			w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return

		} else {
			next(w, r)
		}

	})
}
