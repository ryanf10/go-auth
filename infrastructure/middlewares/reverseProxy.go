package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		parsedUrl, _ := url.Parse(target)
		director := func(req *http.Request) {
			req.Host = parsedUrl.Host
			req.URL = parsedUrl
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
