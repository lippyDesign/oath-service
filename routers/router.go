package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lippyDesign/oath-service.git/repository"
)

// HTTPRouter routes http requests
type HTTPRouter struct {
	User UserRouter
}

// NewRouter initializes HTTP router
func NewRouter(repo *repository.Repo) *HTTPRouter {
	router := gin.Default()
	userRouter := newUserRouter(router, repo)
	r := HTTPRouter{}
	r.User = userRouter
	r.User.RegisterRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return &r
}
