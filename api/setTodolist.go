package main

import (
	"fmt"
	"time"

	"crypto/md5"

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

// Time型をstring型に変換
var timeFormat = "2006-01-02 15:04:05"
func timeToString(t time.Time) string {
    str := t.Format(timeFormat)
    return str
}

func HandleLambdaEvent(event Todo) (error) {
	sess  := session.Must(session.NewSession())
	db    := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("TodoDatabase")

	//ID作成（TitleとTimeで文字列生成）
	strTime := timeToString(time.Now())
	md5 := md5.Sum([]byte(strTime + event.Title))
	id  := fmt.Sprintf("%x", md5)

	w:= Todo{ID:id, Title: event.Title, Num: event.Num, Time: time.Now()}
	err := table.Put(w).Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}