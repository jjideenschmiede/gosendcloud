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
r := gosendcloud.Request{
    PublicKey: "",
    SecretKey: "",
}

// Define request body
body := gosendcloud.CreateAParcelBody{
    Parcel: gosendcloud.CreateAParcelBodyParcel{
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
        Shipment: gosendcloud.CreateAParcelBodyShipment{
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
create, err := gosendcloud.CreateAParcel(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(create)
}
```

### Inserting shipments into an integration

If you want to create an order via an integration, then you can do this as follows. You can find more information about the structure of the request in the [sendcloud documentation](https://docs.sendcloud.sc/api/v2/shipping/#inserting-shipments-into-an-integration).

```go
// Define request
r := gosendcloud.Request{
    PublicKey: "",
    SecretKey: "",
}

// Define request body
var body []gosendcloud.InsertingShipmentsBody

body = append(body, gosendcloud.InsertingShipmentsBody{
    Address:             "Insulindelaan",
    Address2:            "",
    City:                "Eindhoven",
    CompanyName:         "Sendcloud",
    Country:             "NL",
    CreatedAt:           "2021-12-02T10:00:00.555309+00:00",
    Currency:            "EUR",
    CustomsInvoiceNr:    "",
    CustomsShipmentType: nil,
    Email:               "info@jj-development.de",
    ExternalOrderId:     "1234521226",
    ExternalShipmentId:  nil,
    HouseNumber:         "115",
    Name:                "Kwiedor",
    OrderNumber:         "414124124214214",
    OrderStatus: &gosendcloud.InsertingShipmentsBodyOrderStatus{
        Id:      "fulfilled",
        Message: "Fulfilled",
    },
    ParcelItems: []gosendcloud.InsertingShipmentsBodyParcelItems{},
    PaymentStatus: &gosendcloud.InsertingShipmentsBodyPaymentStatus{
        Id:      "paid",
        Message: "Paid",
    },
    PostalCode:                 "5642 CV",
    SenderAddress:              1,
    ShippingMethod:             111,
    ShippingMethodCheckoutName: "DPD Classic",
    Telephone:                  "+31612345678",
    ToPostNumber:               "",
    ToServicePoint:             nil,
    ToState:                    nil,
    TotalOrderValue:            "13.99",
    UpdatedAt:                  "2021-12-02T11:01:47.505309+00:00",
    Weight:                     "0.300",
    Width:                      "40.00",
    Height:                     "30.00",
    Length:                     "50.00",
    Shipments:                  nil,
    CheckoutPayload:            nil,
})

// Add a parcel item
body[0].ParcelItems = append(body[0].ParcelItems, gosendcloud.InsertingShipmentsBodyParcelItems{
    Description:   "T-Shirt",
    HsCode:        "",
    OriginCountry: "",
    ProductId:     "a_random_id",
    Properties:    map[string]interface{}{},
    Quantity:      2,
    Sku:           "a_random_sku",
    Value:         "300",
    Weight:        "0.300",
})

// Inserting a new shipment
insert, err := gosendcloud.InsertingShipments(198416, body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(insert)
}
```