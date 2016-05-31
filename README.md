[![GoDoc](https://godoc.org/github.com/leonelquinteros/gorand?status.svg)](https://godoc.org/github.com/leonelquinteros/gorand)


# Gorand

Goodies for working with random stuff


## ID

Generates a random 64 bytes number represented as a 128 chars string

```go
import "github.com/leonelquinteros/gorand"

func main() {
    value, err := gorand.ID()
    if err != nil {
        panic(err.Error())
    }
    
    println(value)
}
```


## GetHex

Retrieves a fixed amount of bytes hex number represented as a string. 
The result string length will have twice the number of bytes requested. 

```go
import "github.com/leonelquinteros/gorand"

func main() {
    value, err := gorand.GetHex(64)
    if err != nil {
        panic(err.Error())
    }
    
    println(value)
}
```
