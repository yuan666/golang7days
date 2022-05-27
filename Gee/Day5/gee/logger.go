package gee

import (
	"log"
	"time"
)
//log miiddlemare
func Logger() HandlerFunc {
	return func(c *Context) {
		//Start Timer
		t := time.Now()
		//Process request
		c.Next()

		//Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
