package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertificateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertificateResponse", string(data)}, " ")
}
