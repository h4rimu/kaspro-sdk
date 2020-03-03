package utils

import b64 "encoding/base64"

func DecodeBase64(uniqueID string, message string) *string {
	data, err := b64.StdEncoding.DecodeString(message)
	if err != nil {
		log.Errorf(uniqueID, "Error occurred "+err.Error())
		return nil
	}
	result := string(data)
	return &result
}
