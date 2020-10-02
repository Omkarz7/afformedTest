package main

import (
	"affordmed/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	middleware.InitMiddleWare(router)

	serve := &http.Server{
		Addr:         ":4700",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	serve.ListenAndServe()
}
