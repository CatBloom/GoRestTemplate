package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"main/db"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	echoLambda *echoadapter.EchoLambda
	e          *echo.Echo
	database   db.Database
)

func init() {
	log.Println("init")

	database = db.NewDatabase()

	// model

	// controller

	e = echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if len(reqBody) > 0 {
			preJson := bytes.Buffer{}
			if err := json.Indent(&preJson, reqBody, "", "   "); err != nil {
				log.Printf("error:%s", err.Error())
			}
			log.Printf("ReqBody: %s", preJson.String())
		}
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := err.Error()

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message.(string)
		}
		if !c.Response().Committed {
			c.JSON(code, map[string]string{"error": msg})
		}
	}
	echoLambda = echoadapter.New(e)
}

func main() {
	if os.Getenv("ENV") == "local" {
		e.Start(":8080")
	} else {
		lambda.Start(handler)
	}
	defer database.GetDB().Close()
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}
