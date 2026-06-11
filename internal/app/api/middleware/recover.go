package middleware

import (
	"fmt"
	"go-gin-boilerplate/internal/common"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recover(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errRecover *common.APIError
				if apiErr, ok := err.(*common.APIError); ok {
					errRecover = apiErr
				} else if e, ok := err.(error); ok {
					errRecover = &common.APIError{
						Message:   e.Error(),
						ErrorCode: common.ServerError,
						Status:    http.StatusInternalServerError,
						Stack:     string(debug.Stack()),
					}
				} else {
					errRecover = &common.APIError{
						Message:   fmt.Sprintf("%v", err),
						ErrorCode: common.ServerError,
						Status:    http.StatusInternalServerError,
						Stack:     string(debug.Stack()),
					}
				}

				if errRecover.Status >= http.StatusBadRequest {
					logger.WithFields(logrus.Fields{
						"url":    c.Request.URL.Path,
						"method": c.Request.Method,
						"error":  errRecover.Error(),
						"stack":  errRecover.Stack,
					}).Error("Process request error")
				}

				c.AbortWithStatusJSON(errRecover.Status, common.SendError(errRecover))
			}
		}()
		c.Next()
	}
}
