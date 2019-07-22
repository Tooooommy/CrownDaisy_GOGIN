package base_controller

import (
	"CrownDaisy_GOGIN/helper"
	"CrownDaisy_GOGIN/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

type BaseController struct {
}

func (b *BaseController) NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, helper.ReturnResult(helper.CodeNotFoundPage, "page not found", nil))
	return
}

func (b *BaseController) HandleResultPanic(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Debugf("panic occurred: %+v", r)
			logger.Debugln(debug.Stack())
			if _, ok := r.(*helper.Result); ok {
				c.JSON(http.StatusOK, r)
			}
		}
	}()
	c.Next()
}