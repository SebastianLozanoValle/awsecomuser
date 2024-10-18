package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebastianlozanovalle/awsecomuser/awsgo"
	"github.com/sebastianlozanovalle/awsecomuser/db"
	"github.com/sebastianlozanovalle/awsecomuser/models"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.AWSInitialize()

	if !ValidateParameters() {
		fmt.Println("Error en los parámetros. debe enviar 'SecretManager'")
		err := errors.New("Error en los parámetros debe enviar SecretName")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret " + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err

}

func ValidateParameters() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}