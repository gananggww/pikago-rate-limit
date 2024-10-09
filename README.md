# Rate Limit Request

# Installation

## Install
```bash
go get github.com/gananggww/pikago-rate-limit
```

# Usage

## how to inject or call
```go
val := dragonite.NewValidate(rds) //you can inject dep this in main.go in ucase or handler or direct call to func
err = o.val.RateLimit(ctx, dragonite.RateLimitConfig{ 
    SubAs:      "pikachu", // sub type for identifier as who
    In:         10000, // limit request in ms
    As:         "pokemon", // main type for identifier as who  
    Limit:      10, // how many limit request 
}).Error

if err != nil {
    return
}
```
