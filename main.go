package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanilla/go-relationtable/api/config"
	"github.com/vanilla/go-relationtable/api/entities"
	"github.com/vanilla/go-relationtable/api/payload"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/getAllUsers", func(c *gin.Context) {
			var users []entities.User

			db.Debug().Preload("Type").Preload("RoleUsers").Preload("RoleUsers.Role").Preload("RoleUsers.Role.RoleMenus.Menu").Preload("RoleUsers.Role.RoleMenus.Type").Find(&users)

			res := payload.MessageResponse(true, "OK", users)

			c.JSON(http.StatusOK, res)
		})

		api.GET("/getAllRoles", func(c *gin.Context) {
			var roles []entities.Role

			db.Debug().Preload("RoleUsers").Preload("RoleUsers.User").Preload("RoleUsers.Role").Preload("RoleUsers.User.Type").Find(&roles)

			res := payload.MessageResponse(true, "OK", roles)

			c.JSON(http.StatusOK, res)
		})
	}

	r.Run()
}
