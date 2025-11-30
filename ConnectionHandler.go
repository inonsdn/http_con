package http_con

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ConnectionHandler struct {
	route   *gin.Engine
	sigChan chan int
}

func InitHandlerWithGroup(httpPathConfigs []HttpGroupPath) (*ConnectionHandler, error) {
	route := gin.Default()
	var err error
	for _, httpGroupPath := range httpPathConfigs {
		prefix := httpGroupPath.name
		err = registerRoute(route, prefix, httpGroupPath.paths)

		if err != nil {
			break
		}
	}

	return &ConnectionHandler{
		route:   route,
		sigChan: make(chan int, 0),
	}, err
}

func registerRoute(route *gin.Engine, groupPrefix string, paths []HttpPath) error {
	var err error
	for _, httpPath := range paths {

		realPathName := fmt.Sprintf("%s%s", groupPrefix, httpPath.name)

		if httpPath.method == routeMethod_GET {
			route.GET(realPathName, httpPath.callback)
		} else if httpPath.method == routeMethod_POST {
			route.POST(realPathName, httpPath.callback)
		} else {
			err = fmt.Errorf("")
		}
	}
	return err
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
