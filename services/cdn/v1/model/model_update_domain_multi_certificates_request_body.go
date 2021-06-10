package model

import (
	"encoding/json"

	"strings"
)

type UpdateDomainMultiCertificatesRequestBody struct {
	Https *UpdateDomainMultiCertificatesRequestBodyContent `json:"https,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequestBody", string(data)}, " ")
}
