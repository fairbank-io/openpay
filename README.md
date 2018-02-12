# OpenPay API Client

[![Build Status](https://travis-ci.org/fairbank-io/openpay.svg?branch=master)](https://travis-ci.org/fairbank-io/openpay)
[![GoDoc](https://godoc.org/github.com/fairbank-io/openpay?status.svg)](https://godoc.org/github.com/fairbank-io/openpay)
[![Version](https://img.shields.io/github/tag/fairbank-io/openpay.svg)](https://github.com/fairbank-io/openpay/releases)
[![Software License](https://img.shields.io/badge/license-MIT-red.svg)](LICENSE)

Pure Go [OpenPay](https://www.openpay.mx/) client implementation.

## Example

```go
// Start a new client instance
client, _ := openpay.NewClient("API_KEY", "MERCHANT_ID", nil)

// Register customer
rick := &Customer{
    Name:     "Rick",
    LastName: "Sanchez",
    Email:    "rick@mail.com",
    Address:  Address{
        CountryCode: "MX",
        PostalCode:  "94560",
    },
}
client.Customers.Create(rick)

// Add Card
card := &Card{
    HolderName:      "Rick Sanchez",
    CardNumber:      "4111111111111111",
    CVV2:            "401",
    ExpirationMonth: "10",
    ExpirationYear:  "19",
    Address:         rick.Address,
}
client.Charges.AddCard(card)

// Execute charge
sale := &ChargeWithStoredCard{
    Charge: Charge{
        Method:      "card",
        Amount:      1000,
        Currency:    "MXN",
        Description: "sample charge operation",
        Customer:    rick,
    },
    SourceID: card.ID,
    Capture:  true,
}
client.Charges.WithCard(sale)
```
