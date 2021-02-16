package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	handlers "main/handlers"
)

var ginLambda *ginadapter.GinLambda
var app *gin.Engine

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	app := gin.Default()
	app.Static("/img", "./download")
	app.GET("/get", handlers.GetWallpaper)

	ginLambda = ginadapter.New(app)
}

func main() {

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "development" {
		app.Run(":" + os.Getenv("PORT"))
	}

	lambda.Start(Handler)

}
