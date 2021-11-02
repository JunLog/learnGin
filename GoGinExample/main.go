package main

import (
	"fmt"
	"main/pkg/setting"
	"main/routers"
	"net/http"
)

func main() {
	// r := gin.Default()
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "test",
	// 	})
	// })

	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
