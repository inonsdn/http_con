package http_con

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ConnectionHandler struct {
	route   *gin.Engine
	sigChan chan int
}

type RouteRegistration interface {
	RegisterRoute(r *gin.Engine)
}

func NewHandler() *ConnectionHandler {
	route := gin.Default()
	return &ConnectionHandler{
		route:   route,
		sigChan: make(chan int, 0),
	}
}

func (c *ConnectionHandler) RegisterRoute(r RouteRegistration) {
	r.RegisterRoute(c.route)
}

func (c *ConnectionHandler) WaitAndGetStatus() int {
	return <-c.sigChan
}

func (c *ConnectionHandler) Run(addr string) {
	err := c.route.Run(addr)
	if err != nil {
		fmt.Println("Found error")
		c.sigChan <- -1
	}
	c.sigChan <- 0
}
