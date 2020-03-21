package routers

import (
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/lippyDesign/oath-service.git/entities"
	"github.com/lippyDesign/oath-service.git/repository"
)

// UserRouter interface to send/receive user over REST
type UserRouter interface {
	FindOneByID(c *gin.Context)
	FindAll(c *gin.Context)
	UpdateOne(c *gin.Context)
	CreateOne(c *gin.Context)
	DeleteOne(c *gin.Context)
	RegisterRoutes()
}

type userRouter struct {
	repo   *repository.Repo
	engine *gin.Engine
}

// newUserRouter initialize new user REST client
func newUserRouter(engine *gin.Engine, repo *repository.Repo) UserRouter {
	var ur UserRouter = userRouter{repo, engine}
	return ur
}

func (router userRouter) RegisterRoutes() {
	router.engine.GET("/users/:id", router.FindOneByID)
	router.engine.GET("/users", router.FindAll)
	router.engine.POST("/users", router.CreateOne)
	router.engine.PUT("/users/:id", router.UpdateOne)
	router.engine.DELETE("/users/:id", router.DeleteOne)
}

func (router userRouter) FindAll(c *gin.Context) {
	users, err := router.repo.User.FindAll()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, users)
}

func (router userRouter) FindOneByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid user id format"})
		return
	}
	user, err1 := router.repo.User.FindOneByID(intID)
	if err1 != nil {
		c.JSON(400, gin.H{"message": err1.Error()})
		return
	}
	c.JSON(200, user)
}

func (router userRouter) CreateOne(c *gin.Context) {
	var user *entities.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid user payload"})
		return
	}
	err = checkmail.ValidateFormat(user.Email)
	if err != nil {
		c.JSON(400, gin.H{"message": "Unable to create user, Email is invalid"})
		return
	}
	newUser, err := router.repo.User.CreateOne(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func (router userRouter) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid user id format"})
		return
	}
	var user *entities.User
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid user payload"})
		return
	}
	if intID != user.ID {
		c.JSON(400, gin.H{"message": "Unable to update user, IDs cannot be changed"})
		return
	}
	err = checkmail.ValidateFormat(user.Email)
	if err != nil {
		c.JSON(400, gin.H{"message": "Unable to update user, Email is invalid"})
		return
	}
	updatedUser, err1 := router.repo.User.UpdateOne(user)
	if err1 != nil {
		c.JSON(400, gin.H{"message": err1.Error()})
		return
	}
	c.JSON(200, updatedUser)
}

func (router userRouter) DeleteOne(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid user id format"})
		return
	}
	var user *entities.User
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid user payload"})
		return
	}
	if intID != user.ID {
		c.JSON(400, gin.H{"message": "Unable to delete user, IDs do not match"})
		return
	}
	err = router.repo.User.DeleteOne(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.AbortWithStatus(204)
}
