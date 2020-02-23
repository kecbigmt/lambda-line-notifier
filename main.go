package main

import (
  "os"
  "context"

  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"
  "github.com/utahta/go-linenotify"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
  token := os.Getenv("LINE_NOTIFY_ACCESS_TOKEN")
  c := linenotify.New()
  for _, message := range sqsEvent.Records {
    c.Notify(token, message.Body, "", "", nil)
  }
  return nil
}

func main() {
  lambda.Start(handler)
}
