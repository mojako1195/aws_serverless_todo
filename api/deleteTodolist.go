package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Todo struct {
	ID    string	`json:"id"`
}

func HandleLambdaEvent(todo Todo) (error) {
	sess  := session.Must(session.NewSession())
	db    := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("TodoDatabase")

	err   := table.Delete("ID", todo.ID).Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}