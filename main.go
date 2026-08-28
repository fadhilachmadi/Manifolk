package main

import (
	"Manifolk/database"
	"Manifolk/handler"
	"Manifolk/repository"
	"Manifolk/router"
	"fmt"
	"log"
)

func main() {
	db, err := database.InitPostgre()
	if err != nil {
		log.Fatal("Error connecting database:", err)
	}

	defer db.Close()

	todoRepository := repository.NewTodoReposPosgre(db)
	todoHandler := handler.NewTodoHandler(todoRepository)

	r := router.SetupRouter(todoHandler)

	fmt.Println("Server running on http://localhost:8080")

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
