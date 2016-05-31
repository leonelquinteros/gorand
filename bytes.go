package gorand

import (
    "crypto/rand"
)

// GetBytes returns a fixed amount of random bytes. 
// Specify the amount of bytes wanted by passing it on the n parameter.
func GetBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    
    return b
}
