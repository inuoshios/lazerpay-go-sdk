# Lazerpay Golang SDK ⚡

## Installation ✔️
To install the package, use the command below.

```sh
go get github.com/milinches/lazerpay-go-sdk
```

Import the package.

```go
import "github.com/milinches/lazerpay-go-sdk"
```

Optional, import the `godotenv` package in your code. It'll help in loading the `.env` files.

```go
import "github.com/joho/godotenv"
```

You can also get your Test API keys, from the Lazerpay Dashboard [here](https://dashboard.lazerpay.finance/).

`INITIALIZE TRANSACTIONS`

Initiate a payment.

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/milinches/lazerpay-go-sdk"
)

func main() {
    // Load the env file
    if err := godotenv.Load(".Sample.env); err != nil {
        log.Fatal("Error loading .env file")
    }

    secretKey := os.GetEnv("SECRET_KEY")
    publicKey := os.GetEnv("PUBLIC_KEY")

    // You can also handle error by adding a second parameter
    initialize, _ := lazerpay.LazerpayClient(publicKey, secretKey)

    // Initialize payment
    payment, _ := lazerpay.InitTransaction(initialize, "xxxxxxxx", "100", 
    "Njoku Emmanuel", "kalunjoku123@gmail.com", "USD", "USDC", true
    )

    log.Println(payment)
}
```

`VERIFY TRANSACTION`

Helps confirm a transaction after payment has been made.

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/milinches/lazerpay-go-sdk"
)

func main() {
    // Load the env file
    if err := godotenv.Load(".Sample.env); err != nil {
        log.Fatal("Error loading .env file")
    }

    secretKey := os.GetEnv("SECRET_KEY")
    publicKey := os.GetEnv("PUBLIC_KEY")

    // You can also handle error by adding a second parameter
    initialize, _ := lazerpay.LazerpayClient(publicKey, secretKey)

    // Verify Transaction
    verify, _ := lazerpay.VerifyTransaction(initialize, "referenceNumberxxxxx")

    log.Println(verify)
}
```

`Get Accepted Coins`

Returns the list of accepted coins at [Lazerpay](https://www.lazerpay.finance/).

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/milinches/lazerpay-go-sdk"
)

func main() {
    // Load the env file
    if err := godotenv.Load(".Sample.env); err != nil {
        log.Fatal("Error loading .env file")
    }

    secretKey := os.GetEnv("SECRET_KEY")
    publicKey := os.GetEnv("PUBLIC_KEY")

    // You can also handle error by adding a second parameter
    initialize, _ := lazerpay.LazerpayClient(publicKey, secretKey)

    // Get the list of all accepted coins
    coins, _ := lazerpay.GetAcceptedCoins(initialize)

    log.Println(coins)
}
```

`Transfer Funds`

Helps send funds to another recipient address.

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/milinches/lazerpay-go-sdk"
)

func main() {
    // Load the env file
    if err := godotenv.Load(".Sample.env); err != nil {
        log.Fatal("Error loading .env file")
    }

    secretKey := os.GetEnv("SECRET_KEY")
    publicKey := os.GetEnv("PUBLIC_KEY")

    // You can also handle error by adding a second parameter
    initialize, _ := lazerpay.LazerpayClient(publicKey, secretKey)

    // Send funds to another recipient address
    transfer, _ := lazerpay.Transfer(initialize, 10, "recipientxxxxxx", "USDC", "Binance Smart Chain")

    log.Println(transfer)
}
```