package middlewares

import (
	"EffectiveMobile/init/logger"
	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/server/http/handlers"
	"EffectiveMobile/pkg/constants"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func ValidateBody() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dbPeople = new(entities.People)
		var bodyBytes []byte

		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)
		}

		if len(bodyBytes) == 0 {
			handlers.NewErrorResponse(ctx, constants.FailedParseBody)
			return
		}

		err := json.Unmarshal(bodyBytes, dbPeople)
		if err != nil {
			logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

			handlers.NewErrorResponse(ctx, err)
			return
		}

		if dbPeople.Name == "" {
			handlers.NewErrorResponse(ctx, constants.InvalidName)
			return
		}

		if dbPeople.Surname == "" {
			handlers.NewErrorResponse(ctx, constants.InvalidSurname)
			return
		}

		if dbPeople.PassportSeries <= 999 || dbPeople.PassportSeries >= 10000 {
			handlers.NewErrorResponse(ctx, constants.InvalidSeries)
			return
		}

		if dbPeople.PassportNumber <= 99999 || dbPeople.PassportNumber >= 1000000 {
			handlers.NewErrorResponse(ctx, constants.InvalidNumber)
			return
		}

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		ctx.Next()
	}
}
