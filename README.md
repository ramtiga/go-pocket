# go-pocket

Interface to pocket API written by golang.

## Install:

    go get github.com/ramtiga/go-pocket

## Usage:

    import (
            "fmt"
            "github.com/ramtiga/go-pocket"
            "log"
    )

    func main() {
            client := pocket.NewClient("CONSUMER_KEY", "ACCESS_TOKEN")

            opt := map[string]interface{}{
                    "Search": "iphone",
                    "Count":  10,
            }

            results, err := client.PocketList(opt)
            if err != nil {
                    log.Fatal(err)
            }

            for _, res := range results {
                    fmt.Println(res.Resolved_title)
            }
    }

## Author:

ramtiga
