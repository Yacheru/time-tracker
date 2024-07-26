package tasks

import (
	"EffectiveMobile/internal/entities"
	"EffectiveMobile/pkg/constants"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) ActiveTaskExists(ctx *gin.Context, id int) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT * FROM tasks WHERE people_id = $1 AND end_task IS NULL
		)
	`
	err := t.db.GetContext(ctx.Request.Context(), &exists, query, id)

	return exists, err
}

func (t *TaskRepository) DeleteTask(ctx *gin.Context, id int) (*entities.Task, error) {
	var dbTask = new(entities.Task)

	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}

	delQuery := `
		DELETE FROM tasks WHERE people_id = $1 RETURNING *
	`
	err = tx.GetContext(ctx.Request.Context(), dbTask, delQuery, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	updateQuery := `
		UPDATE peoples SET task_id = null WHERE id = $1
	`
	_, err = tx.ExecContext(ctx.Request.Context(), updateQuery, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return dbTask, tx.Commit()
}

func (t *TaskRepository) StartTask(ctx *gin.Context, id int) (*entities.Task, error) {
	var dbTask = new(entities.Task)

	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}

	exist, err := t.ActiveTaskExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, constants.HaveActiveTask
	}

	taskQuery := `
		INSERT INTO tasks (people_id) VALUES ($1) RETURNING *;
	`
	err = tx.GetContext(ctx.Request.Context(), dbTask, taskQuery, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	peopleQuery := `
		UPDATE peoples SET task_id=$1 WHERE id = $2
	`
	_, err = tx.ExecContext(ctx, peopleQuery, dbTask.ID, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return dbTask, tx.Commit()
}

func (t *TaskRepository) StopTask(ctx *gin.Context, id int) (*entities.Task, error) {
	var dbTask = new(entities.Task)
	now := time.Now().Unix()

	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	exist, err := t.ActiveTaskExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, constants.NoActiveTask
	}

	taskQuery := `
		UPDATE tasks SET end_task = $1, labor = $2 - start_task WHERE people_id = $3 AND end_task IS NULL RETURNING *;
	`
	err = tx.GetContext(ctx.Request.Context(), dbTask, taskQuery, now, now, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	peopleQuery := `
		UPDATE peoples SET task_id=NULL WHERE id = $1
	`
	_, err = tx.ExecContext(ctx.Request.Context(), peopleQuery, id)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return dbTask, tx.Commit()
}

func (t *TaskRepository) GetAllTasks(ctx *gin.Context, id, limit int) (*[]entities.Task, error) {
	var dbTasks = new([]entities.Task)

	query := `
		SELECT * FROM tasks WHERE people_id = $1 ORDER BY labor DESC LIMIT $2
	`

	err := t.db.SelectContext(ctx.Request.Context(), dbTasks, query, id, limit)
	if err != nil {
		return nil, err
	}

	return dbTasks, nil
}
