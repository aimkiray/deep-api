package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/aimkiray/deep-api/pkg/setting"
)

func main() {
	router := gin.Default()
	router.GET("/moon/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.ParseFloat(id, 64)
		if err == nil && len(id) >= 2 {
			file, err := os.Open("akari233.yaml")
			if err != nil {
				return
			}
			defer file.Close()
			content, err := ioutil.ReadAll(file)

			contentStr := string(content)

			contentStr = strings.Replace(contentStr, "def_str", "re_str"+id[1:2], -1)

			c.Writer.WriteHeader(http.StatusOK)
			c.Header("Content-Disposition", "attachment; filename=akari233.yaml")
			c.Header("Content-Type", "application/text/plain")
			c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
			c.Writer.Write([]byte(contentStr))
		}
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
