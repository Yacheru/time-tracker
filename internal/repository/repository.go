package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/repository/postgres/peoples"
	"EffectiveMobile/internal/repository/postgres/tasks"
)

type PeopleRepository interface {
	DeletePeople(ctx *gin.Context, passportSeries, passportNumber int) (*entities.People, error)
	PeopleExists(ctx *gin.Context, passportSeries, passportNumber int) (bool, error)
	GetPeople(ctx *gin.Context, passportSeries, passportNumber int) (*entities.People, error)
	CreatePeople(ctx *gin.Context, people *entities.People) (*entities.People, error)
	GetAllPeoples(ctx *gin.Context, limit int) (*[]entities.People, error)
}

type TaskRepository interface {
	DeleteTask(ctx *gin.Context, id int) (*entities.Task, error)
	ActiveTaskExists(ctx *gin.Context, id int) (bool, error)
	//GetTask(ctx *gin.Context, p *entities.People) (*entities.Task, error)
	GetAllTasks(ctx *gin.Context, id, limit int) (*[]entities.Task, error)
	StartTask(ctx *gin.Context, id int) (*entities.Task, error)
	StopTask(ctx *gin.Context, id int) (*entities.Task, error)
}

type Repository struct {
	PeopleRepository
	TaskRepository
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PeopleRepository: peoples.NewPeopleRepository(db),
		TaskRepository:   tasks.NewTaskRepository(db),
	}
}
