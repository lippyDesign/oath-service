package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lippyDesign/oath-service.git/repository"
	"github.com/lippyDesign/oath-service.git/routers"
)

// App is application manager.
type App struct {
	Router *routers.HTTPRouter
	Repo   *repository.Repo
}

func main() {
	a := App{}
	// repo
	newRepo, err := repository.NewRepo()
	if err != nil {
		log.Println(err)
		return
	}
	a.Repo = newRepo
	defer a.Repo.DB.Close()
	// router
	a.Router = routers.NewRouter(a.Repo)
}
