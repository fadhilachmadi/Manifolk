package main

import (
	"context"
	"log"

	"Manifolk/database"
	todoHandler "Manifolk/handler"
	"Manifolk/repository"
	"Manifolk/router"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {

	db, err := database.InitPostgre()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTodoReposPosgre(db)
	todoHandler := todoHandler.NewTodoHandler(repo)

	r := router.SetupRouter(todoHandler)

	ginLambda = ginadapter.New(r)
}

func lambdaHandler(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {

	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(lambdaHandler)
}
