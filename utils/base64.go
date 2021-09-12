package utils

import "encoding/base64"

func ToBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func FromBase64(text string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
