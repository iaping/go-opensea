# go-opensea
Golang sdk for https://docs.opensea.io/reference/api-overview, Implement the v2 version. Not all interfaces implemented, just a basic framework, you can implement other interfaces to submit to me.

## Installation

go get：
```shell
go get github.com/iaping/go-opensea/v2
```

## Examples
[examples](examples/main.go)

## Simple
```go
package main

import (
	"fmt"

	"github.com/iaping/go-opensea/v2"
)

func main() {
    client := opensea.New("your key", nil)
	resp, err := client.GetListings().Do("matic", "seaport", &opensea.ListingsParam{
		AssetContractAddress: "0xad154cc15761b47ccb8d0d21114f855134b9a01d",
		Token_ids:            []int{871},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

## Help
Let me know if you have any problems.