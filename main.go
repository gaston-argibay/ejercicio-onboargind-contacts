package main

import (
	"context"
	"ejercicio/domain"
	"ejercicio/repository"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Printf("REQUEST ID: %s", lc.AwsRequestID)
	log.Printf("FUNCTION NAME: %s", lambdacontext.FunctionName)

	contact := domain.Contact{Id: "2",
		FirstName: "juana",
		LastName:  "arco",
		Status:    "CREATED"}

	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	client := dynamodb.New(sess)

	err := repository.AddTableItem(client, contact)
	if err != nil {
		return fmt.Sprintf("Got an error adding item to table:"), err
	}
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
