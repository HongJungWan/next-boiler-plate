package router

import (
	"context"
	"eCommerce/config"
	"eCommerce/repository"
	"eCommerce/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Router struct {
	config *config.Config

	engin *gin.Engine

	service    *service.Service
	repository *repository.Repository
}

func NewRouter(config *config.Config, service *service.Service, repository *repository.Repository) (*Router, error) {
	r := &Router{
		config:     config,
		engin:      gin.New(),
		service:    service,
		repository: repository,
	}

	// 서버 실행 코드
	r.engin.Use(requestTimeOutMiddleWare(1 * time.Second))

	NewMongoRouter(r, r.service.MService)
	NewMySQLRouter(r, r.service.MySQLService)

	return r, r.engin.Run(config.ServerInfo.Port)
}

func requestTimeOutMiddleWare(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			defer close(done)
			c.Next()
		}()

		select {
		case <-done:
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{"error": "Request Timeout"})
				return
			}
		}

	}
}
