package gorand

import ()

// ID is a wrapper for GetHex(64).
// 64 bytes should be enough to use as unique IDs to avoid collisions between recently generated values.
func ID() (string, error) {
	return GetHex(64)
}

// UUID generates a version 4 (randomly generated) UUID as defined in RFC 4122.
// It returns the canonical string representation "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx".
func UUID() (string, error) {
	// Get initial random bytes
	uuid, err := GetHex(16)
	if err != nil {
		return "", err
	}

	// Get random variant
	v, err := GetRandomChars("89ab", 1)
	if err != nil {
		return "", err
	}

	return uuid[0:8] + "-" + uuid[8:12] + "-4" + uuid[13:16] + "-" + v + uuid[17:20] + "-" + uuid[20:], nil
}
