package basehandler

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetDocs(c *gin.Context) {
	// redirect to swagger page
	location := url.URL{Path: "/docs/index.html"}
	c.Redirect(http.StatusFound, location.RequestURI())
}
