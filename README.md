[![GoDoc](https://godoc.org/github.com/leonelquinteros/gorand?status.svg)](https://godoc.org/github.com/leonelquinteros/gorand)


# Gorand

Goodies for working with random stuff.

Package gorand defines a set of utility methods to work with random data. 

The main goal is to curate a collection of functions for random data generation 
in different formats to be used for different purposes. 

It uses the "crypto/rand" package to provide the most secure random data, 
unless indicated otherwise where performance can be the main focus of the method. 

Most implementations are really trivial and anybody could write the same on their own packages, 
but here we can centralize all of them and keep a unified way of retrieving random data. 

Unified QA is another motivator to have and use this package.

Below are some examples of functions inside this package, for the entire reference and docs please refer to the documentation at (https://godoc.org/github.com/leonelquinteros/gorand)


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


## GetAlphaNumString

Returns a fixed length string of random letters and numbers [a-z][A-Z][0-9]

```go
import "github.com/leonelquinteros/gorand"

func main() {
    value, err := gorand.GetAlphaNumString(24)
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
