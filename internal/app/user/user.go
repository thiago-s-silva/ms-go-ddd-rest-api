package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"gorm.io/gorm"
)

var (
	repository *UserRepository
	db         *gorm.DB
	service    *UserService
)

type Dependencies struct {
	Config       *config.Config
	V1RouteGroup *gin.RouterGroup
	Db           *gorm.DB
}

func Load(d *Dependencies) error {
	// init db
	db = d.Db

	// auto migrate User table
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}

	// init user repository to entire User package
	repository = &UserRepository{}

	// init user service to entire User package
	service = &UserService{}

	// register v1 routes
	registerV1Routes(d.V1RouteGroup)

	return nil
}

func registerV1Routes(v1 *gin.RouterGroup) {
	// get all users
	v1.GET("/users", GetAllHandler)

	// create one user
	v1.POST("/user", CreateHandler)

	// get one user by id
	v1.GET("/user/:id", GetOneHandler)

	// delete user by id
	v1.DELETE("/user/:id", DeleteHandler)

	// update user by id
	v1.PUT("/user/:id", UpdateHandler)
}
