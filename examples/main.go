package main

import (
	"fmt"

	"github.com/iaping/go-opensea/v2"
)

var client = opensea.New("your key", nil)

func main() {
	GetListings()
}

func GetListings() {
	resp, err := client.GetListings().Do("matic", "seaport", &opensea.ListingsParam{
		AssetContractAddress: "0xad154cc15761b47ccb8d0d21114f855134b9a01d",
		Token_ids:            []int{871},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
