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

// @Summary      Create People
// @Description  Create people
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        body body entities.People true "Body"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /peoples/create [post]
func (h *Handlers) CreatePeople(ctx *gin.Context) {
	var bodyPeople = new(entities.People)

	res, err := ctx.GetRawData()
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	err = json.Unmarshal(res, bodyPeople)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	people, err := h.peopleService.CreatePeople(ctx, bodyPeople)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "New people added successfully", people)
}

// @Summary      Get All Peoples
// @Description  Get All Peoples
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        limit query int false "Limit"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /peoples [get]
func (h *Handlers) GetAllPeoples(ctx *gin.Context) {
	bodyPeople, err := h.peopleService.GetAllPeoples(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully get all peoples", bodyPeople)
}

// @Summary      Delete People
// @Description  Delete People
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        passportSeries   query int true "Passport series"
// @Param        passportNumber   query int true "Passport number"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /peoples/delete [delete]
func (h *Handlers) DeletePeople(ctx *gin.Context) {
	bodyPeople, err := h.peopleService.DeletePeople(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully delete people", bodyPeople)
}

// @Summary      Update People
// @Description  Update People
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        body body entities.People true "Body"
// @Param        passportSeries   query int true "Passport series"
// @Param        passportNumber   query int true "Passport number"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /peoples/update [patch]
func (h *Handlers) UpdatePeople(ctx *gin.Context) {
	var bodyPeople = new(entities.People)

	res, err := ctx.GetRawData()
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return
	}

	err = json.Unmarshal(res, bodyPeople)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return
	}

	people, err := h.peopleService.UpdatePeople(ctx, bodyPeople)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully update people", people)
}

// @Summary      Start Task
// @Description  Start Task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        passportSeries   query int true "Passport series"
// @Param        passportNumber   query int true "Passport number"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /tasks/start [post]
func (h *Handlers) StartTask(ctx *gin.Context) {
	bodyTask, err := h.taskService.StartTask(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "New task added successfully", bodyTask)
}

// @Summary      Stop Task
// @Description  Stop Task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        passportSeries   query int true "Passport series"
// @Param        passportNumber   query int true "Passport number"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /tasks/stop [post]
func (h *Handlers) StopTask(ctx *gin.Context) {
	bodyTask, err := h.taskService.StopTask(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return

	}

	NewSuccessResponse(ctx, http.StatusOK, "Task stopped successfully", bodyTask)
}

// @Summary      Get All Tasks
// @Description  Get All Tasks
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        passportSeries   query int true "Passport series"
// @Param        passportNumber   query int true "Passport number"
// @Param        limit   path int false "Limit"
// @Success      200  {object}  handlers.Response
// @Failure      400  {object}  handlers.Response
// @Failure      404  {object}  handlers.Response
// @Failure      500  {object}  handlers.Response
// @Router       /tasks/ [get]
func (h *Handlers) GetAllTasks(ctx *gin.Context) {
	bodyTask, err := h.taskService.GetAllTasks(ctx)
	if err != nil {
		logger.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Handler, constants.LoggerURL: ctx.Request.URL.String()})

		NewErrorResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Successfully get all tasks", bodyTask)
}
