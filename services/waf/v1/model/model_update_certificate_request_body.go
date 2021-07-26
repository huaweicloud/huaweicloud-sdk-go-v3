package model

import (
	"encoding/json"

	"strings"
)

type UpdateCertificateRequestBody struct {
	// 证书名

	Name *string `json:"name,omitempty"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}
