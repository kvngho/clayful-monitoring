package command

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/kvngho/clayful-monitoring/internal/infra"
	"github.com/kvngho/clayful-monitoring/internal/pkg/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	SyncCmd.Flags()
}

var SyncCmd = &cobra.Command{
	Use: "sync",
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())
		viper.AutomaticEnv()
		log.Println("sync start")

		sqs := infra.NewSQSService(infra.SQSConfigFromViper())
		
		messageRecv := &infra.SQSMessage{}
		messageList := make([]infra.SQSMessage, 0, 100)
		for {
			messages, err := sqs.Receive()
			if len(messages.Messages) == 0 {
				break
			}
			if err != nil {
				log.Fatalln("error: ", err)
			}
			for _, message := range messages.Messages {
				go func(message types.Message) {
					defer sqs.ExplicitACK(message.ReceiptHandle)
				}(message)

				fmt.Println(*message.Body)
				if err := json.Unmarshal([]byte(*message.Body), messageRecv); err != nil {
					fmt.Println("error unmarshal")
				}
				messageList = append(messageList, *messageRecv)
				sqs.ExplicitACK(message.ReceiptHandle)
			}
		}
		slack.SendSlackNotification(viper.GetString("webhook_url"), messageList)
		return nil
	},
}