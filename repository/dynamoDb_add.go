package repository

import (
	"ejercicio/domain"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func AddTableItem(svc dynamodbiface.DynamoDBAPI, contact domain.Contact) error {
	table := "argibay-ejercicio-contacts"
	av, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return err
	}
	fmt.Printf("marshalled struct: %+v", av)

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: &table,
	})
	if err != nil {
		return err
	}

	return nil
}
