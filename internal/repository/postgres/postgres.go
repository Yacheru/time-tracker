package postgres

import (
	"context"
	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"EffectiveMobile/init/logger"
	"EffectiveMobile/pkg/constants"
)

func NewPostgresConnection(ctx context.Context, dsn string) (*sqlx.DB, error) {
	logger.Info("Open postgresql connection...", logrus.Fields{constants.LoggerCategory: constants.Postgres})
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		logger.ErrorF("Failed open postgresql connection... %v", logrus.Fields{constants.LoggerCategory: constants.Postgres}, err)

		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		logger.ErrorF("Failed ping postgresql... %v", logrus.Fields{constants.LoggerCategory: constants.Postgres}, err)

		return nil, err
	}

	logger.Info("Successful connection to postgres", logrus.Fields{constants.LoggerCategory: constants.Postgres})

	if err := goose.UpContext(ctx, db.DB, "./schema"); err != nil {
		logger.ErrorF("Error using migrates: %v", logrus.Fields{constants.LoggerCategory: constants.Postgres}, err)
	}

	return db, nil
}
