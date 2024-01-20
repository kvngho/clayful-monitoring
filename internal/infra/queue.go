package infra

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/spf13/viper"
)

type Queue interface {
	Publish(ctx context.Context, body string) error
}

type SQSConfig struct {
	URL string
}

// func SQSBindPFlags() *pflag.FlagSet {
// 	f := pflag.NewFlagSet("sqs", pflag.ContinueOnError)
// 	f.String("url", "", "sqs url")
// 	f.String("region", "", "region")
// 	f.String("SecretKey")
// }

type SQS struct {
	Svc *sqs.Client
	Url string
}

func SQSConfigFromViper() SQSConfig {
	return SQSConfig{
		URL: viper.GetString("URL"),
	}
}

func NewSQSService(cfg SQSConfig) *SQS {
	aws_cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalln("sqs error")
	}
	return &SQS{
		Svc: sqs.NewFromConfig(aws_cfg),
		Url: cfg.URL,
	}
}

func (s *SQS) Publish(message SQSMessage) *sqs.SendMessageOutput{
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	stringJson := string(jsonMsg)
	input := &sqs.SendMessageInput{
		QueueUrl: &s.Url,
		MessageBody: &stringJson,
	}
	output, err := s.Svc.SendMessage(context.Background(), input)
	if err != nil {
		log.Fatalln("send message error")
		return nil
	}
	return output
}

func (s *SQS) Receive() (*sqs.ReceiveMessageOutput, error){
	input := &sqs.ReceiveMessageInput{
		QueueUrl: &s.Url,
		MaxNumberOfMessages: *aws.Int32(10),
		WaitTimeSeconds: *aws.Int32(20),
	}

	resp, err := s.Svc.ReceiveMessage(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQS) ExplicitACK (receiptHandle *string) error {
	_, err := s.Svc.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
		QueueUrl: &s.Url,
		ReceiptHandle: receiptHandle,
	})
	return err
}

type SQSMessage struct {
	Type string `json:"type"`
	ClayfulID string `json:"clayful_id"`
}