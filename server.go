package todo

import "net/http"

type Server struct {
	httpServer *http.Server
}

func(s *Server) Run(post string, handler http.Handler) error {
	s.http = &http.Server{
		Addr:			   ":" + port,
		Handler:		   handler,
		MaxHeaderBytes:    1 << 20,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second, 
	}

	return s.httpServer.ListenandServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}