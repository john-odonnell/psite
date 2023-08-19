package web

import (
	"fmt"
	"io/fs"
	"net/http"

	v1 "github.com/john-odonnell/psite/v2/pkg/api/v1"
)

// TODO: Target API
//   /
//   /cv
//   /hobbies
//   /links
//   /messageboard
//   /projects
//   /projects/euler
//   /projects/urlshort

// Server represents an instance of the webserver being run on a given port,
// with the given mapping of API endpoints to http.Handler functions.
type Server struct {
	port   int
	routes map[string]http.Handler
}

// NewServer returns a Server instance given a port and a filesystem of static
// resources.
func NewServer(staticDir fs.FS, port int) Server {
	return Server{
		port: port,
		routes: map[string]http.Handler{
			"/":        v1.DefaultRootHandler(),
			"/static/": http.FileServer(http.FS(staticDir)),
		},
	}
}

// Listen serves the configured API endpoints and http.Handlers on the
// configured port.
func (s Server) Listen() error {
	mux := http.NewServeMux()
	for endpoint, handler := range s.routes {
		mux.Handle(endpoint, handler)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), mux)
}
