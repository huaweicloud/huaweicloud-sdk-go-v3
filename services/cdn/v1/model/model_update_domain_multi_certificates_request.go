package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainMultiCertificatesRequest struct {
	Body *UpdateDomainMultiCertificatesRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequest", string(data)}, " ")
}
