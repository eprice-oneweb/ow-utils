package basehandler

import (
	"embed"
	"net/http"
	"strings"

	"github.com/eprice-oneweb/ow-utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
)

var FileSystem embed.FS

// @Summary     Gets application info
// @Description Gets application info
// @Tags        info
// @Accept      */*
// @Produce     json
// @Success     200 {object} map[string]string
// @Router      /info [get]
func Info(commitHash string) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		data, err := FileSystem.ReadFile("VERSION")
		if err != nil {
			logger.Logger.Errorf("An error occurred when reading the VERSION file %v", err)
		}

		m := make(map[string]string)
		for _, e := range strings.Fields(string(data)) {
			parts := strings.Split(e, "=")
			m[strcase.ToLowerCamel(strings.ToLower(parts[0]))] = parts[1]
		}

		m["commitHash"] = commitHash

		c.JSON(http.StatusOK, m)
	}

	return gin.HandlerFunc(fn)
}
