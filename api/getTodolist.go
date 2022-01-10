package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Todo struct {
	ID    string	`dynamo:"ID"`
	Title string 	`dynamo:"Title" json:"title"`
	Num   string	`dynamo:"Num"   json:"num"`
	Time  time.Time
}

func HandleLambdaEvent(event Todo) ([]Todo,error) {
	sess  := session.Must(session.NewSession())
	db    := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("TodoDatabase")

	// get all items
	var results []Todo
	err := table.Scan().All(&results)
	if err != nil {
		return results,err
	}

	return results,nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}