package middlewares

import (
	"EffectiveMobile/internal/server/http/handlers"
	"EffectiveMobile/pkg/constants"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ValidateParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		passportSeries, err := strconv.Atoi(ctx.Query("passportSeries"))
		if err != nil {
			handlers.NewErrorResponse(ctx, constants.FailedValidateParams)
			return
		}

		passportNumber, err := strconv.Atoi(ctx.Query("passportNumber"))
		if err != nil {
			handlers.NewErrorResponse(ctx, constants.FailedValidateParams)
			return
		}

		if passportSeries <= 999 || passportSeries >= 10000 {
			handlers.NewErrorResponse(ctx, constants.InvalidSeries)
			return
		}

		if passportNumber <= 99999 || passportNumber >= 1000000 {
			handlers.NewErrorResponse(ctx, constants.InvalidNumber)
			return
		}

		ctx.Next()
	}
}
