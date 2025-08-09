package router

import (
	// "fmt"

	"github.com/LuizFernando991/golang-api/api/controllers"
	// "github.com/LuizFernando991/golang-api/infra/config"
	// "github.com/LuizFernando991/golang-api/infra/database"
)

type Controllers struct {
	UserController *controllers.UserController
}

func GetControllers() Controllers {
	// envVariables := config.GetEnvConfig()
	// databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", envVariables.DB_USER, envVariables.DB_PASSWORD, envVariables.DB_HOST, envVariables.DB_PORT, envVariables.DB_DATABASE)
	// dbConn := database.GetDBConn(databaseUrl)

	UserControler := controllers.NewUserController()

	return Controllers{
		UserController: UserControler,
	}
}
