package main

import (
    "golang_api/routes"
)

func main() {
    router := routes.InitRouter()
    router.Run()
}
