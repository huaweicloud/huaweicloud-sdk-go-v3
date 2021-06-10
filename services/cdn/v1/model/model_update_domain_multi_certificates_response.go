package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDomainMultiCertificatesResponse struct {
	Https          *UpdateDomainMultiCertificatesResponseBodyContent `json:"https,omitempty"`
	HttpStatusCode int                                               `json:"-"`
}

func (o UpdateDomainMultiCertificatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesResponse", string(data)}, " ")
}
