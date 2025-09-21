
package middleware

import "net/http"

func SetupMiddleware(mux *http.ServeMux) *http.ServeMux {
    // Add any middleware logic here
    return mux
}
