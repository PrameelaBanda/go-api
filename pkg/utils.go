package pkg

import (
	b64 "encoding/base64"
	"os"
)

func B64Encode(data []byte) string {
	return b64.StdEncoding.EncodeToString(data)
}

func B64Decode(data string) ([]byte, error) {
	b, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetEnvOrDefault(envVarName, defaultVal string) string {
	v := os.Getenv(envVarName)
	if v == "" {
		v = defaultVal
	}

	return v
}


