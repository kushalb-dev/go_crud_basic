package migrate

import (
	"github.com/kushalb-dev/go_crud/initializers"
	"github.com/kushalb-dev/go_crud/models"
)

func init() {
	// initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}

func Migrate() {
	initializers.DB.AutoMigrate(&models.Post{})
}
