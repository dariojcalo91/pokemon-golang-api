package adapters

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	s := &Server{}

	s.routes()
	return s
}

// Run attaches the router and starts listening and serving HTTP requests
func (s *Server) Run(addr ...string) error {
	return s.router.Run(addr...)
}
