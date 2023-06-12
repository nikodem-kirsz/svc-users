package app

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	userspb "github.com/nikodem-kirsz/svc-users/api/proto/users"
	"github.com/nikodem-kirsz/svc-users/app/configuration"
	"github.com/nikodem-kirsz/svc-users/app/db/model"
	"github.com/nikodem-kirsz/svc-users/app/handler"
)

type App struct {
	Log   *logrus.Entry
	Conf  *configuration.Config
	Model *model.MysqlStorage
	// GormDriver *gorm.DB
	Handler userspb.UsersServer
}

func BuildDependencies(conf *configuration.Config, log *logrus.Entry) (*App, error) {
	// creating msql model and orchestrator
	m := model.NewMySqlStorage()

	// creating grpc handler provider with injected db instance model
	h, err := handler.New(m)
	if err != nil {
		errors.Wrap(err, "failed to create default handler validators")
	}
	// Creating fully injected App service to be exposed on grpc server entity
	return &App{
		Log:  log,
		Conf: conf,
		// GormDriver: gormDriver,
		Handler: h,
	}, nil
}
