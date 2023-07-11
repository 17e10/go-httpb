# go-httpb

[![GoDev][godev-image]][godev-url]

go-httpb パッケージは HTTP client のユーティリティを提供します.

## Usage

```go
import "github.com/17e10/go-httpb"

resp, err := http.Get("https://go.dev/")
if resp.StatusCode != http.StatusOK {
    return httpb.ErrStatus(resp)
}
```

## License

This software is released under the MIT License, see LICENSE.

## Author

17e10

[godev-image]: https://pkg.go.dev/badge/github.com/17e10/go-httpb
[godev-url]: https://pkg.go.dev/github.com/17e10/go-httpb
