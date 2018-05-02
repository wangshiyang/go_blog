package test

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.GET("/bin/search", func(c *gin.Context) {

		time.Sleep(time.Millisecond * 1)

		c.String(http.StatusOK, "OK")
	})
	s := &http.Server{
		Addr:           ":1234",
		Handler:        route,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
