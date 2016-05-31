package gorand

import ()

// ID is a wrapper for GetHex(64).
// 64 bytes should be enough to use as unique IDs to avoid collisions between recently generated values. 
func ID() (string, error) {
    return GetHex(64)
}
