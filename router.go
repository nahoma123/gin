package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Router data will be registered to http listener
type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

type routing struct {
	host           string
	domain         string
	allowedOrigins string
	routers        []Router
}

// Routers contains the functions of http handler to clean payloads and pass it the service
type Routers interface {
	Serve()
}

// Initialize is for initialize the handler
func Initialize(host, allowedOrigins string, routers []Router, domain string) Routers {
	return &routing{
		host,
		domain,
		allowedOrigins,
		routers,
	}
}

// Serve is to start serving the HTTP Listener for every domain
func (r *routing) Serve() {
	server := gin.Default()

	for _, router := range r.routers {

		if len(router.Middlewares) != 0 {
			// Append the router to the middlware
			router.Middlewares = append(router.Middlewares, router.Handler)
			server.Handle(router.Method, router.Path, router.Middlewares...)
		} else {
			server.Handle(router.Method, router.Path, router.Handler)

		}

	}

	logrus.WithFields(logrus.Fields{
		"host":   r.host,
		"domain": r.domain,
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(r.host, server))

}
