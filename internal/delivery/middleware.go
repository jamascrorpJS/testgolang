package delivery

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamascrorpJS/eBank/pkg/cryptography"
)

func CheckHeader() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.GetHeader("X-UserId")
		digest := c.GetHeader("X-Digest")
		c.Set("id", id)
		c.Set("digest", digest)
		if id == "" && digest == "" {
			c.AbortWithError(401, errors.New("want header"))
		}
	}
}

func Digest(c *gin.Context) {
	digest := c.MustGet("digest").(string)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	validate, err := cryptography.Validate(body, digest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	if !validate {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "digest",
		})
		c.Abort()
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	c.Next()
}
