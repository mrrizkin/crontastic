package system

import (
	"github.com/mrrizkin/crontastic/routes"

	"github.com/mrrizkin/crontastic/app/config"
	"github.com/mrrizkin/crontastic/app/models"
	"github.com/mrrizkin/crontastic/system/database"
	"github.com/mrrizkin/crontastic/system/server"
	"github.com/mrrizkin/crontastic/system/session"
	"github.com/mrrizkin/crontastic/system/types"
	"github.com/mrrizkin/crontastic/third_party/logger"
	"github.com/mrrizkin/crontastic/third_party/whatsapp"
)

func Run() {
	conf := config.New()
	log := logger.New(conf)
	sess := session.New(conf)
	model := models.New()
	db := database.New(conf, model, log)
	wa := whatsapp.New(conf)
	defer db.Stop()
	err := db.Start()
	if err != nil {
		panic(err)
	}

	serv := server.New(conf, log, sess)

	routes.Setup(&types.App{
		App:      serv.App,
		Logger:   log,
		Database: db,
		Config:   conf,
		Session:  sess,
		WhatsApp: wa,
	})

	err = serv.Serve()
	if err != nil {
		panic(err)
	}
}
