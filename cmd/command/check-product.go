package command

import (
	"context"
	"fmt"
	"log"

	"github.com/kvngho/clayful-monitoring/internal/infra"
	"github.com/kvngho/clayful-monitoring/internal/pkg/utils"
	"github.com/kvngho/clayful-monitoring/internal/userinterface"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	CheckProductCmd.Flags()
}

var CheckProductCmd = &cobra.Command{
	Use: "check-product",
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())
		viper.AutomaticEnv()
		log.Println("check-product start")

		client := infra.NewEntCilent(infra.EntConfigFromENV())
		defer client.DB.Close()

		clayful := infra.NewClayful(infra.ClayfulConfigFromENV())

		sqs := infra.NewSQSService(infra.SQSConfigFromViper())
		checker := userinterface.NewChecker(client, clayful)
		products, err := checker.FindProductDeepingAndClayful(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		for _, product := range products {
			if product.Clayful_id != "" {
				fmt.Printf("Checking::%s::%s\n", product.Name, product.Clayful_id)
				res := checker.FindProductFromClayful(context.Background(), product.Clayful_id)
				switch product.Clayful_options.(type) {
				case string:
					if string(product.Clayful_options.(string)) != res.Variants[0].ID {
						sqs.Publish(infra.SQSMessage{
							Type: "product-clayfulid",
							ClayfulID: product.Clayful_id,
						})
						fmt.Println("clayful id is not matched")
					}

					if product.Price_deeping != res.Price.Sale.Raw {
						sqs.Publish(infra.SQSMessage{
							Type: "product-price",
							ClayfulID: product.Clayful_id,
						})
						fmt.Println("price is not matched")
					}
				default:
					setDeepingOption := make(map[string]struct{})
					setClayfulOption := make(map[string]struct{})

					optionMap, _ := product.Clayful_options.(map[string]interface{})
					for _, v := range(optionMap) {
						setDeepingOption[v.(string)] = struct{}{}
					}
					for _, v := range(res.Variants) {
						setClayfulOption[v.ID] = struct{}{}
						if product.Price_deeping != v.Price.Sale.Raw {
							sqs.Publish(infra.SQSMessage{
								Type: "product-price",
								ClayfulID: product.Clayful_id,
							})
							fmt.Println("price is not matched")
						}
					}
					if !utils.IsEqual(setClayfulOption, setDeepingOption) {
						sqs.Publish(infra.SQSMessage{
							Type: "product-options",
							ClayfulID: product.Clayful_id,
						})
					}
				}
				
			}
		}
		return nil
	},
}