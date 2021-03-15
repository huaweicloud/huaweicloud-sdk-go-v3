package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckCertificateResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckCertificateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckCertificateResponse struct{}"
	}

	return strings.Join([]string{"CheckCertificateResponse", string(data)}, " ")
}
