package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kvngho/clayful-monitoring/internal/infra"
	"github.com/kvngho/clayful-monitoring/internal/userinterface"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	CheckCouponCmd.Flags()
}

var CheckCouponCmd = &cobra.Command{
	Use: "check-coupon",
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())
		viper.AutomaticEnv()
		log.Println("check-coupon start")

		client := infra.NewEntCilent(infra.EntConfigFromENV())
		defer client.DB.Close()

		clayful := infra.NewClayful(infra.ClayfulConfigFromENV())

		sqs := infra.NewSQSService(infra.SQSConfigFromViper())
		checker := userinterface.NewChecker(client, clayful)
		coupons, err := checker.FindAvailableCouponDeeping(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		for _, coupon := range coupons {
			if coupon.ID != 0 {
				fmt.Printf("Checking::%s::%s\n", coupon.Name, coupon.Clayful_id)
				res := checker.FindCouponFromClayful(context.Background(), coupon.Clayful_id)
				deepingCouponTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", coupon.EndDate)
				if err != nil {
					log.Fatalln("error: ", err)
				}
				clayfulCouponTime, _ := time.Parse("2006-01-02T15:04:05.000Z", res.ExpiresAt.Raw)
				if err != nil {
					log.Fatalln("error: ", err)
				}
				if !deepingCouponTime.Equal(clayfulCouponTime) {
					sqs.Publish(infra.SQSMessage{
						Type: "coupon",
						ClayfulID: coupon.Clayful_id,
					})
				}
			}
		}
		return nil
	},
}