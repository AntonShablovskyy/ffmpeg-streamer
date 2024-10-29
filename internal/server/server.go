package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	port   int
	hlsDir string
}

func New(port int, hlsDir string) *Server {
	return &Server{
		port:   port,
		hlsDir: hlsDir,
	}
}

func (s *Server) ServeHls() {
	http.Handle("/", s.addHeaders(http.FileServer(http.Dir(s.hlsDir))))

	fmt.Printf("Serving HLS files from directory '%s' on port %d\n", s.hlsDir, s.port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.port), nil))
}

func (s *Server) addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
