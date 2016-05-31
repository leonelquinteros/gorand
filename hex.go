package gorand

import (
    "fmt"
)

// GetHex returns a string representation of a (n) bytes random hexadecimal number. 
// The length of the result string will be twice as long as the amount of bytes requested due to hex representation. 
GetHex(n int) (string, error) {
    b, err := GetBytes(n)
    if err != nil {
        return "", err.Error()
    }
    
    return fmt.Sprintf("%x", b), nil
}
