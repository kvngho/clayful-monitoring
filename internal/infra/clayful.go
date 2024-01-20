package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kvngho/clayful-monitoring/internal/pkg/clayful"
	"github.com/spf13/viper"
)

func ClayfulConfigFromENV() ClayfulConfig {
	return ClayfulConfig{
		TOKEN: viper.GetString("CLAYFUL_TOKEN"),
	
	}
}

type ClayfulConfig struct {
	TOKEN 		string `env:"CLAYFUL_TOKEN"`
}

var CLAYFUL_HEADER map[string]string = map[string]string{
	"Accept": "application/json",
	// "Accept-Encoding": "gzip",
	"Content-Type": "application/json",
	"Authorization": fmt.Sprintf("%s %s", "Bearer", viper.GetString("CLAYFUL_TOKEN")),
}

type Clayful struct {
	header map[string]string
}

func NewClayful(cfg ClayfulConfig) Clayful {
	return Clayful{
		header: map[string]string{
			"Accept": "application/json",
			"Content-Type": "application/json",
			"Authorization": fmt.Sprintf("%s %s", "Bearer", cfg.TOKEN),
		},
	}
}

func (c *Clayful) GetProduct(ctx context.Context, productID string) *clayful.ClayfulProduct{
	url := fmt.Sprintf("https://api.clayful.io/v1/products/%s", productID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalln("req error")
		return nil
	}

	for k, v := range c.header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("do error")
		return nil
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read error")
		return nil
	}

	clayfulResponse := &clayful.ClayfulProduct{}

	if err := json.Unmarshal(body, clayfulResponse); err != nil {
		log.Fatalln("Unmarshal error")
		return nil
	}
	return clayfulResponse
}

func (c *Clayful) GetCoupon(ctx context.Context, couponID string) *clayful.ClayfulCoupon{
	url := fmt.Sprintf("https://api.clayful.io/v1/coupons/%s", couponID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalln("req error")
		return nil
	}

	for k, v := range c.header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("do error")
		return nil
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read error")
		return nil
	}

	clayfulResponse := &clayful.ClayfulCoupon{}

	if err := json.Unmarshal(body, clayfulResponse); err != nil {
		log.Fatalln("Unmarshal error")
		return nil
	}
	return clayfulResponse
}