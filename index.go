package main

import (
    "golang_api/routes"
    "gorm.io/gorm"
    "golang_api/config"
)

var (
    db             *gorm.DB                  = config.SetupDatabaseConnection()
)

func main() {
    defer config.CloseDatabaseConnection(db)
    router := routes.InitRouter()
    router.Run()
}
