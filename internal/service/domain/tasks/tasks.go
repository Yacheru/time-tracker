package tasks

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"EffectiveMobile/init/logger"
	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/repository"
	"EffectiveMobile/pkg/constants"
)

type TaskService struct {
	tasks   repository.TaskRepository
	peoples repository.PeopleRepository
}

func NewTaskService(tasks repository.TaskRepository, peoples repository.PeopleRepository) *TaskService {
	return &TaskService{
		tasks:   tasks,
		peoples: peoples,
	}
}

func (t *TaskService) StartTask(ctx *gin.Context) (*entities.Task, error) {
	passportSeries, _ := strconv.Atoi(ctx.Query("passportSeries"))
	passportNumber, _ := strconv.Atoi(ctx.Query("passportNumber"))

	p, err := t.peoples.GetPeople(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("GetPeople: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	dbTask, err := t.tasks.StartTask(ctx, *p.ID)
	if err != nil {
		switch {
		case errors.Is(err, constants.HaveActiveTask):
			return nil, constants.HaveActiveTask
		default:
			logger.DebugF("StartTask: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

			return nil, constants.ErrStartTask
		}
	}

	return dbTask, nil
}
func (t *TaskService) StopTask(ctx *gin.Context) (*entities.Task, error) {
	passportSeries, _ := strconv.Atoi(ctx.Query("passportSeries"))
	passportNumber, _ := strconv.Atoi(ctx.Query("passportNumber"))

	exists, err := t.peoples.PeopleExists(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("PeopleExists: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}
	if !exists {
		return nil, constants.ErrPeopleNotFound
	}

	people, err := t.peoples.GetPeople(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("GetPeople: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	dbTask, err := t.tasks.StopTask(ctx, *people.ID)
	if err != nil {
		switch {
		case errors.Is(err, constants.NoActiveTask):
			return nil, constants.NoActiveTask
		default:
			logger.DebugF("StopTask: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

			return nil, constants.ErrStopTask
		}
	}

	return dbTask, nil
}

func (t *TaskService) GetAllTasks(ctx *gin.Context) (*[]entities.Task, error) {
	passportSeries, _ := strconv.Atoi(ctx.Query("passportSeries"))
	passportNumber, _ := strconv.Atoi(ctx.Query("passportNumber"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}

	exists, err := t.peoples.PeopleExists(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("PeopleExists: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}
	if !exists {
		return nil, constants.ErrPeopleNotFound
	}

	people, err := t.peoples.GetPeople(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("GetPeople: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	dbTasks, err := t.tasks.GetAllTasks(ctx, *people.ID, limit)
	if err != nil {
		logger.DebugF("GetAllTasks: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, constants.ErrGetAllTasks
	}

	return dbTasks, nil
}
