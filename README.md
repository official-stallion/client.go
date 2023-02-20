<h1 align="center">
    Stallion Golang SDK
</h1>

<br />

Golang client SDK for Stallion message broker.

## How to use?

Get package:

```shell
go get github.com/official-stallion/go-sdk@latest
```

### Creating Clients

You can connect to stallion server like the example below:

```go
package main

import (
	"fmt"
	"time"

	sdk "github.com/official-stallion/go-sdk"
)

func main() {
	client, err := sdk.NewClient("st://localhost:9090")
	if err != nil {
		panic(err)
	}
}
```

### Subscribe on a topic

```go
client.Subscribe("topic", func(data []byte) {
    // any handler that you want
    fmt.Println(string(data))
})
```

### Publish over a topic

```go
client.Publish("topic", []byte("Hello"))
```

### Unsubscribe from a topic

```go
client.Unsubscribe("topic")
```

## Connect with Auth

Connect with username and password set in url:

```go
client, err := stallion.NewClient("st://root:Pa$$word@localhost:9090")
if err != nil {
    panic(err)
}
```

## Mock

You can create a mock client to create a stallion client sample:

```go
package main

import sdk "github.com/official-stallion/go-sdk"

func main() {
	client := sdk.NewMockClient()
	
	client.Publish("topic", []byte("message"))
}
```
