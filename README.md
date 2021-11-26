# gosendcloud

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gosendcloud.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gosendcloud/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gosendcloud/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gosendcloud)](https://goreportcard.com/report/github.com/jjideenschmiede/gosendcloud) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gosendcloud?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gosendcloud) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gosendcloud) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

With this library it should be possible to interact with the endpoints of the sendcloud API via golang functions. Since we can not implement this library project-specific completely, you are welcome to extend it with us.

## Install

```console
go get github.com/jjideenschmiede/gosendcloud
```

## How to use?

### Create a Parcel

If a new dispatch note is to be created, then this can be done directly with the following function. You can find more information about the structure of the request in the [sendcloud documentation](https://docs.sendcloud.sc/api/v2/shipping/#create-a-parcel).

```go
// Define request
r := Request{
    publicKey: "",
    secretKey: "",
}

// Define request body
body := CreateAParcelBody{
    Parcel: CreateAParcelBodyParcel{
        Name:         "Jonas Kwiedor",
        CompanyName:  "J&J Ideenschmiede GmbH",
        Address:      "Mercatorstraße",
        HouseNumber:  "32a",
        City:         "Geesthacht",
        PostalCode:   "21502",
        Telephone:    "+4941528903730",
        RequestLabel: true,
        Email:        "info@jj-ideenschmiede.de",
        Country:      "DE",
        Shipment: CreateAParcelBodyShipment{
            Id: 8,
        },
        Weight:                     "10.000",
        OrderNumber:                "41267142142",
        InsuredValue:               2000,
        TotalOrderValueCurrency:    "EUR",
        TotalOrderValue:            "24.99",
        Quantity:                   1,
        ShippingMethodCheckoutName: "DPD",
    },
}

// Create a new parcel
create, err := CreateAParcel(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(create)
}
```