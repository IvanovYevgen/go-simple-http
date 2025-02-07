package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

type assetsResponce struct {
	Assets    []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

// type assetResponce struct {
// 	Asset     assetData `json:"data"`
// 	Timestamp int64     `json:"timestamp"`
// }

func (d assetData) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s [NAME] %s [PRICE]", d.ID, d.Rank, d.Name, d.PriceUsd)
}

type assetData struct {
	ID            string `json:"id"`
	Rank          string `json:"rank"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Supply        string `json:"supply"`
	MaxSupply     string `json:"max"`
	MarketCapUsd  string `json:"marketCapUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd      string `json:"priceUsd"`
}

func (l loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return l.next.RoundTrip(req)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("Redirected from:", via[0].URL)
			return nil
		},
		Transport: loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Response body:", string(body))
	var r assetsResponce
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	for _, v := range r.Assets {
		fmt.Println(v.Info())
	}
}
