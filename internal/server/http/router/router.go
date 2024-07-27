package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	_ "EffectiveMobile/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"EffectiveMobile/internal/repository"
	"EffectiveMobile/internal/server/http/handlers"
	mw "EffectiveMobile/internal/server/http/middlewares"
	"EffectiveMobile/internal/service"
)

type Routes struct {
	handlers *handlers.Handlers
	router   *gin.RouterGroup
}

func NewRoutes(router *gin.RouterGroup, db *sqlx.DB) *Routes {
	repo := repository.NewPostgresRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewHandler(services.PeopleService, services.TaskService)

	return &Routes{
		handlers: handler,
		router:   router,
	}
}

func (r *Routes) Routers() {
	peoples := r.router.Group("/peoples")
	{
		peoples.GET("/", r.handlers.GetAllPeoples)
		peoples.POST("/create", mw.ValidateBody(), r.handlers.CreatePeople)
		peoples.DELETE("/delete", mw.ValidateParams(), r.handlers.DeletePeople)
		peoples.PATCH("/update", mw.ValidateParams(), mw.ValidateBody(), r.handlers.UpdatePeople)
	}
	tasks := r.router.Group("/tasks")
	{
		tasks.GET("/", mw.ValidateParams(), r.handlers.GetAllTasks)
		tasks.POST("/start", mw.ValidateParams(), r.handlers.StartTask)
		tasks.POST("/stop", mw.ValidateParams(), r.handlers.StopTask)
	}

	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
