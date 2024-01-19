package infra

import (
	"context"
	"log"

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

func (s *SQS) Publish(messageBody string) *sqs.SendMessageOutput{
	input := &sqs.SendMessageInput{
		QueueUrl: &s.Url,
		MessageBody: &messageBody,
	}
	output, err := s.Svc.SendMessage(context.Background(), input)
	if err != nil {
		log.Fatalln("send message error")
		return nil
	}
	return output
}