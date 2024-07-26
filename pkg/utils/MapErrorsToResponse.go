package utils

import (
	"EffectiveMobile/pkg/constants"
	"database/sql"
	"errors"
	"net/http"
)

func MapErrorsToResponse(err error) (int, string) {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound, constants.ErrPeopleNotFound.Error()
	case errors.Is(err, constants.ErrPeopleNotFound):
		return http.StatusNotFound, constants.ErrPeopleNotFound.Error()
	case errors.Is(err, constants.FailedParseBody):
		return http.StatusBadRequest, constants.FailedParseBody.Error()
	case errors.Is(err, constants.ErrPeopleExist):
		return http.StatusConflict, constants.ErrPeopleExist.Error()
	case errors.Is(err, constants.InvalidNumber):
		return http.StatusBadRequest, constants.InvalidNumber.Error()
	case errors.Is(err, constants.InvalidSeries):
		return http.StatusBadRequest, constants.InvalidSeries.Error()
	case errors.Is(err, constants.InvalidSeries):
		return http.StatusInternalServerError, constants.InvalidSeries.Error()
	case errors.Is(err, constants.ErrStopTask):
		return http.StatusInternalServerError, constants.ErrStopTask.Error()
	case errors.Is(err, constants.ErrGetAllTasks):
		return http.StatusInternalServerError, constants.ErrGetAllTasks.Error()
	case errors.Is(err, constants.NoActiveTask):
		return http.StatusNotFound, constants.NoActiveTask.Error()
	case errors.Is(err, constants.HaveActiveTask):
		return http.StatusConflict, constants.HaveActiveTask.Error()
	case errors.Is(err, constants.FailedValidateParams):
		return http.StatusBadRequest, constants.FailedValidateParams.Error()
	case errors.Is(err, constants.InvalidSurname):
		return http.StatusBadRequest, constants.InvalidSurname.Error()
	case errors.Is(err, constants.InvalidName):
		return http.StatusBadRequest, constants.InvalidName.Error()
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}
