package middlewares

import (
	"EffectiveMobile/internal/server/http/handlers"
	"EffectiveMobile/pkg/constants"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ValidateParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := strconv.Atoi(ctx.Query("passportSeries"))
		if err != nil {
			handlers.NewErrorResponse(ctx, constants.FailedValidateParams)
			return
		}

		_, err = strconv.Atoi(ctx.Query("passportNumber"))
		if err != nil {
			handlers.NewErrorResponse(ctx, constants.FailedValidateParams)
			return
		}

		ctx.Next()
	}
}
