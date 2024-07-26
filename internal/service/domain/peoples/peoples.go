package peoples

import (
	"EffectiveMobile/init/logger"
	"EffectiveMobile/internal/entities"
	"EffectiveMobile/internal/repository"
	"EffectiveMobile/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type PeopleService struct {
	people repository.PeopleRepository
}

func NewPeopleService(people repository.PeopleRepository) *PeopleService {
	return &PeopleService{
		people: people,
	}
}

func (ps *PeopleService) DeletePeople(ctx *gin.Context) (*entities.People, error) {
	passportSeries, _ := strconv.Atoi(ctx.Query("passportSeries"))
	passportNumber, _ := strconv.Atoi(ctx.Query("passportNumber"))

	dbPeople, err := ps.people.DeletePeople(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("DeletePeople: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	return dbPeople, nil
}

func (ps *PeopleService) PeopleExists(ctx *gin.Context, passportSeries int, passportNumber int) (bool, error) {
	exists, err := ps.people.PeopleExists(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("PeopleExists: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return false, err
	}

	return exists, nil
}

func (ps *PeopleService) GetPeople(ctx *gin.Context, passportSeries int, passportNumber int) (*entities.People, error) {
	exist, err := ps.people.PeopleExists(ctx, passportNumber, passportSeries)
	if err != nil {
		logger.DebugF("PeopleExists: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}
	if !exist {
		return nil, constants.ErrPeopleNotFound
	}

	dbPeople, err := ps.people.GetPeople(ctx, passportSeries, passportNumber)
	if err != nil {
		logger.DebugF("GetPeople: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	return dbPeople, nil
}

func (ps *PeopleService) GetAllPeoples(ctx *gin.Context) (*[]entities.People, error) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}

	dbPeoples, err := ps.people.GetAllPeoples(ctx, limit)
	if err != nil {
		logger.DebugF("GetAllPeoples: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	return dbPeoples, nil
}

func (ps *PeopleService) CreatePeople(ctx *gin.Context, people *entities.People) (*entities.People, error) {
	dbPeople, err := ps.people.CreatePeople(ctx, people)
	if err != nil {
		logger.DebugF("CreatePeople: %s", logrus.Fields{constants.LoggerCategory: constants.Service}, err.Error())

		return nil, err
	}

	return dbPeople, nil
}
