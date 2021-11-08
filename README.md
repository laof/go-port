## goport

Web server function, get IP and Port

#### Install

Install the package with:
```
go get github.com/laof/goport
```

Import it with:
```
import "github.com/laof/goport"
```

#### Example
```go
ip := goport.GetIp()
port := goport.InputPort("8918")
```