package main

import (
	"fmt"
	"github.com/IvanovYevgen/http-project/coincap"
	"log"
	"time"
)

func main() {
	coincapClient, err := coincap.NewClient(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}
	// assets, err := coincapClient.GetAssets()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, asset := range assets {
	// 	fmt.Println(asset.Info())
	// }
	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitcoin.Info())

}
