package main

import (
	"github.com/AnibrataMalkhandi/Google-Map-Leads-Generation/config"
	"github.com/AnibrataMalkhandi/Google-Map-Leads-Generation/db"
	"github.com/AnibrataMalkhandi/Google-Map-Leads-Generation/pkg/api/seed"
	"github.com/AnibrataMalkhandi/Google-Map-Leads-Generation/pkg/server"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		seed.Migrate(dbHandler)
	}

	server.Setup(dbHandler)
}
