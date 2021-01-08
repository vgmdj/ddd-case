package httpfacade

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vgmdj/ddd-case/interfaces/httpfacade/router"
)

// HTTPServer gin server
type HTTPServer struct {
	srv  *gin.Engine
	port string
}

// HTTPConfig config of http
type HTTPConfig struct {
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

// NewGinServer return the object of http facade
func NewGinServer(c *HTTPConfig) *HTTPServer {
	port := strings.Trim(c.Port, ":")

	return &HTTPServer{
		srv:  gin.New(),
		port: fmt.Sprintf(":%s", port),
	}
}

// Launch http server run
func (hf *HTTPServer) Launch(ctx context.Context) error {
	router.Register(hf.srv)

	errChan := make(chan error)
	go func() {
		err := hf.srv.Run(hf.port)
		errChan <- err
	}()

	select {
	case <-ctx.Done():
		log.Println("server exit due to other context")
		return nil

	case err := <-errChan:
		return err

	}
}
