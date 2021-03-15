package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type PushCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PushCertificateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PushCertificateResponse struct{}"
	}

	return strings.Join([]string{"PushCertificateResponse", string(data)}, " ")
}
