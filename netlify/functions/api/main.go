package main

import (
	"encoding/json"
	"net/http"

	"Manifolk/database"
	"Manifolk/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(
	request events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {

	db, err := database.InitPostgre()
	if err != nil {
		return response(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	defer db.Close()

	repo := repository.NewTodoReposPosgre(db)

	todos, err := repo.GetAllData()
	if err != nil {
		return response(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return response(
		http.StatusOK,
		map[string]interface{}{
			"message": "Get all todo list",
			"todos":   todos,
		},
	)
}

func response(
	statusCode int,
	data interface{},
) (*events.APIGatewayProxyResponse, error) {

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
