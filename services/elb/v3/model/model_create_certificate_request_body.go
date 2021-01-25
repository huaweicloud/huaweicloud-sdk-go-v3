package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateCertificateRequestBody struct {
	Certificate *CreateCertificateOption `json:"certificate"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}
