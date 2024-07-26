package handlers

import (
	"EffectiveMobile/init/logger"
	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/service"
	"EffectiveMobile/pkg/constants"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handlers struct {
	peopleService service.PeopleService
	taskService   service.TaskService
}

func NewHandler(peopleService service.PeopleService, taskService service.TaskService) *Handlers {
	return &Handlers{
		peopleService: peopleService,
		taskService:   taskService,
	}
}

func (h *Handlers) CreatePeople(ctx *gin.Context) {
	var dbPeople = new(entities.People)

	res, err := ctx.GetRawData()
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	err = json.Unmarshal(res, dbPeople)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	dbPeople, err = h.peopleService.CreatePeople(ctx, dbPeople)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "New people added successfully", dbPeople)
}

func (h *Handlers) StartTask(ctx *gin.Context) {
	dbTask, err := h.taskService.StartTask(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "New task added successfully", dbTask)
}

func (h *Handlers) StopTask(ctx *gin.Context) {
	dbTask, err := h.taskService.StopTask(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return

	}

	NewSuccessResponse(ctx, http.StatusOK, "Task stopped successfully", dbTask)
}

func (h *Handlers) GetAllTasks(ctx *gin.Context) {
	dbTask, err := h.taskService.GetAllTasks(ctx)
	if err != nil {
		logger.ErrorF("%v: %s", logrus.Fields{constants.LoggerCategory: constants.Handler}, ctx.Request.URL, err.Error())

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully get all tasks", dbTask)
}

func (h *Handlers) GetAllPeoples(ctx *gin.Context) {
	dbPeople, err := h.peopleService.GetAllPeoples(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully get all peoples", dbPeople)
}

func (h *Handlers) DeletePeople(ctx *gin.Context) {
	dbPeople, err := h.peopleService.DeletePeople(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully delete people", dbPeople)
}
