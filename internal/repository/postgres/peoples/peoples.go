package peoples

import (
	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/repository/postgres/tasks"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type PeopleRepository struct {
	db    *sqlx.DB
	tasks *tasks.TaskRepository
}

func NewPeopleRepository(db *sqlx.DB) *PeopleRepository {
	taskRepo := tasks.NewTaskRepository(db)

	return &PeopleRepository{
		db:    db,
		tasks: taskRepo,
	}
}

func (p *PeopleRepository) PeopleExists(ctx *gin.Context, passportSeries, passportNumber int) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1 FROM peoples WHERE passport_series=$1 OR passport_number=$2
		)
	`
	err := p.db.GetContext(ctx.Request.Context(), &exists, query, passportSeries, passportNumber)

	return exists, err
}

func (p *PeopleRepository) DeletePeople(ctx *gin.Context, passportSeries, passportNumber int) (*entities.People, error) {
	var dbPeople = new(entities.People)

	query := `
		DELETE FROM peoples
		WHERE passport_series=$1 OR passport_number=$2
		RETURNING id, surname, name, patronymic, passport_series, passport_number, task_id 
	`

	err := p.db.GetContext(ctx.Request.Context(), dbPeople, query, passportSeries, passportNumber)
	if err != nil {
		return nil, err
	}

	return dbPeople, nil
}

func (p *PeopleRepository) GetAllPeoples(ctx *gin.Context, limit int) (*[]entities.People, error) {
	var dbPeoples = new([]entities.People)

	query := `
		SELECT id, surname, name, patronymic, passport_series, passport_number, task_id FROM peoples LIMIT $1
	`

	err := p.db.SelectContext(ctx.Request.Context(), dbPeoples, query, limit)
	if err != nil {
		return nil, err
	}

	return dbPeoples, nil
}

func (p *PeopleRepository) GetPeople(ctx *gin.Context, passportSeries, passportNumber int) (*entities.People, error) {
	var dbPeople = new(entities.People)

	query := `
		SELECT id, surname, name, patronymic, passport_series, passport_number, task_id 
		FROM peoples 
		WHERE passport_series = $1 AND passport_number = $2
	`

	err := p.db.GetContext(ctx.Request.Context(), dbPeople, query, passportSeries, passportNumber)
	if err != nil {
		return nil, err
	}

	return dbPeople, nil
}

func (p *PeopleRepository) CreatePeople(ctx *gin.Context, people *entities.People) (*entities.People, error) {
	var dbPeople = new(entities.People)

	query := `
		INSERT INTO peoples (surname, name, patronymic, passport_series, passport_number) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, surname, name, patronymic, passport_series, passport_number, task_id;
	`
	err := p.db.GetContext(ctx.Request.Context(), dbPeople, query, people.Surname, people.Name, people.Patronymic, people.PassportSeries, people.PassportNumber)
	if err != nil {
		return nil, err
	}

	return dbPeople, nil
}

func (p *PeopleRepository) UpdatePeople(ctx *gin.Context, people *entities.People, id int) (*entities.People, error) {
	var dbPeople = new(entities.People)

	query := `
		UPDATE peoples 
		SET surname=$1, name=$2, patronymic=$3, passport_series=$4, passport_number=$5 
		WHERE id=$6 
		RETURNING id, surname, name, patronymic, passport_series, passport_number, task_id ;
	`
	err := p.db.GetContext(ctx.Request.Context(), dbPeople, query,
		people.Surname, people.Name, people.Patronymic, people.PassportSeries, people.PassportNumber, id,
	)
	if err != nil {
		return nil, err
	}

	return dbPeople, nil
}
