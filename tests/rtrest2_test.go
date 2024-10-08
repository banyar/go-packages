package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/banyar/go-packages/pkg/adapters"
	"github.com/banyar/go-packages/pkg/common"
)

func TestRest2(t *testing.T) {
	ticketId := 3924319
	url := os.Getenv("API_URL")
	apiUrl := fmt.Sprintf("%s/ticket/%d", url, ticketId)
	apiToken := os.Getenv("API_TOKEN")
	httpAdapter := adapters.NewHttpAdapter(apiUrl, apiToken)
	fmt.Println("apiUrl", httpAdapter.BaseURL)
	fmt.Println("apiToken", httpAdapter.Token)

	resp, err := httpAdapter.HttpService.Get()
	if err != nil {
		log.Fatal("ERROR : ", err)
	}
	common.DisplayJsonFormat("rest2Adapter", resp)
}
