package main

import (
    "log"
    "todo-app/database"
    "todo-app/routes"
)

func main() {
    // config.LoadEnv()
    database.InitDatabase()

    r := routes.SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
