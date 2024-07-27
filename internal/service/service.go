package service

import (
	"github.com/gin-gonic/gin"

	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/repository"
	"EffectiveMobile/internal/service/domain/peoples"
	"EffectiveMobile/internal/service/domain/tasks"
)

type PeopleService interface {
	DeletePeople(ctx *gin.Context) (*entities.People, error)
	UpdatePeople(ctx *gin.Context, people *entities.People) (*entities.People, error)
	GetAllPeoples(ctx *gin.Context) (*[]entities.People, error)
	PeopleExists(ctx *gin.Context, passportSeries int, passportNumber int) (bool, error)
	GetPeople(ctx *gin.Context, passportSeries int, passportNumber int) (*entities.People, error)
	CreatePeople(ctx *gin.Context, people *entities.People) (*entities.People, error)
}

type TaskService interface {
	GetAllTasks(ctx *gin.Context) (*[]entities.Task, error)
	StartTask(ctx *gin.Context) (*entities.Task, error)
	StopTask(ctx *gin.Context) (*entities.Task, error)
}

type Service struct {
	TaskService
	PeopleService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TaskService:   tasks.NewTaskService(r.TaskRepository, r.PeopleRepository),
		PeopleService: peoples.NewPeopleService(r.PeopleRepository),
	}
}
