# go-pocket

Interface to pocket API written by golang.

## Install:

    go get github.com/ramtiga/go-pocket

## Usage:

    client := pocket.NewClient("CONSUMER_KEY", "ACCESS_TOKEN")

    results, err := client.PocketList()
    if err != nil {
            log.Fatal(err)
    }

    for _, res := range results {
            fmt.Println(res.Resolved_title)
    }


## Author:

ramtiga
