package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kvngho/clayful-monitoring/internal/infra"
)

type SlackRequestBody struct {
	Blocks []SlackMessage `json:"blocks"`
}

type SlackMessage struct {
	Type string `json:"type"`
	Text *SlackMessageText `json:"text,omitempty"`
	Elements []*SlackMessageText`json:"elements,omitempty"`
}

type SlackMessageText struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Emoji bool `json:"emoji,omitempty"`
}

func SendSlackNotification(webhookUrl string, msg []infra.SQSMessage) error {
	slackBody, _ := json.Marshal(MakeSlackMessageStruct(msg))
	fmt.Println(string(slackBody))
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func MakeSlackMessageStruct(msg []infra.SQSMessage) SlackRequestBody {
	var finalMessage SlackRequestBody
	header := SlackMessage{
		Type: "header",
		Text: &SlackMessageText{
			Type: "plain_text",
			Text: "👀 Clayful Sync Error Report",
			Emoji: true,
		},
	}
	divider := SlackMessage{
		Type: "divider",
	}

	timeTable := SlackMessage{
		Type: "section",
		Text: &SlackMessageText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("%s 기준", time.Now().Format("2006-01-02 15:04:05")),
		},
	}
	var syncErrors = make([]SlackMessage, 0, 100)
	var syncErrorsString string
	for _, _msg := range msg {
		syncErrorsString = fmt.Sprintf("%s\n\n %s -- %s\n\n", syncErrorsString, _msg.Type, _msg.ClayfulID)
		// syncErrors = append(syncErrors, SlackMessage{
		// 	Type:     "context",
		// 	Elements: []*SlackMessageText{
		// 		{
		// 			Type: "mrkdwn",
		// 			Text: fmt.Sprintf("%s -- %s", _msg.Type, _msg.ClayfulID),
		// 		},
		// 	},
		// })
	}
	if syncErrorsString != "" {
		syncErrors = append(syncErrors, SlackMessage{
			Type:     "context",
			Elements: []*SlackMessageText{
				{
					Type: "mrkdwn",
					Text: syncErrorsString,
				},
			},
		})
	} else {
		syncErrors = append(syncErrors, SlackMessage{
			Type: "section",
			Text: &SlackMessageText{
				Type: "mrkdwn",
				Text: "깔끔합니다..",
			},
		})
	}
	


	finalMessage.Blocks = append(finalMessage.Blocks, header, divider, timeTable, divider)
	finalMessage.Blocks = append(finalMessage.Blocks, syncErrors...)
	return finalMessage
}