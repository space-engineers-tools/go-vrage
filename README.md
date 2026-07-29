# go-vrage

> [!CAUTION]
> **Active Development Notice**: This project is in an early stage. APIs are unstable and **breaking changes** may be introduced without prior notice. Use in production environments is strongly discouraged. When the first stable release is published, this notice will be removed.

<!-- start of badges -->
![License](https://img.shields.io/github/license/space-engineers-tools/go-vrage)
![Version](https://img.shields.io/github/v/tag/space-engineers-tools/go-vrage?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/space-engineers-tools/go-vrage.svg)](https://pkg.go.dev/github.com/space-engineers-tools/go-vrage)
<!-- end of badges -->

[go-vrage](https://github.com/space-engineers-tools/go-vrage) is a Go client for the VRage Remote API of [Space Engineers 1](https://www.spaceengineersgame.com).

It can be used to programmatically manage and monitor [Dedicated Servers](https://www.spaceengineersgame.com/dedicated-servers), including retrieving server status, managing players, and executing server commands.

## Features

- **Abstracted HMAC Authentication:** Automatically manages nonces, timestamps, and HMAC-SHA1 signature generation.
- **Dual Response Format:** Delivers strongly typed Go structs while preserving raw JSON payload bytes for custom parsing or logging.
- **Idiomatic Error Handling:** Maps API responses into custom Go error types for straightforward error inspection.
- **Context Support:** Full `context.Context` propagation for timeout control and request cancellation across all endpoints.
- **Customizable HTTP Client:** Easily inject custom `http.Client` instances or middleware configurations.

## Installation

To install the package to your Go module, run the following command:

```bash
go get -u github.com/space-engineers-tools/go-vrage
```

To use the package in your code, import it as follows:

```go
import "github.com/space-engineers-tools/go-vrage"
```

## Usage

> [!WARNING]
> **Security Notice**: Never hardcode your security key in source code. Use environment variables for serious applications. The example below is for demonstration purposes only.

<!-- todo -->
```go
package main

import (
    "fmt"
    "time"

    "github.com/space-engineers-tools/go-vrage"
)

func main() {
    // Create a new VRage client with your server configuration
    client := vrage.NewClient(vrage.ClientConfig{
        BaseURL:     "your-base-url",
        SecurityKey: "your-security-key",
    })

    status, err := client.Server()
    if err != nil {
        fmt.Println("Error retrieving server status:", err)
        return
    }
    fmt.Println("Server status:", status)
}
```

## Contributing

We welcome contributions! Please read our [contributing guidelines](CONTRIBUTING.md) for details on how to get started.

## Disclaimer

This project is not affiliated with [Keen Software House](https://www.keenswh.com) or the [Space Engineers](https://www.spaceengineersgame.com) game.
