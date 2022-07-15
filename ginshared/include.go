package ginshared

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Included []string

func GetIncludedParams(c *gin.Context) (included Included) {
	rawInclude := c.Query("include")
	if len(rawInclude) > 0 {
		included = strings.Split(rawInclude, ",")
	}

	return
}

func (i Included) IsIncluded(param string) bool {
	for _, val := range i {
		if val == param {
			return true
		}
	}

	return false
}
