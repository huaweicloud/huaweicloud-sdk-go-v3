package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDomainPermissionsForAgencyRequest struct {
	DomainId string `json:"domain_id"`

	AgencyId string `json:"agency_id"`
}

func (o ListDomainPermissionsForAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDomainPermissionsForAgencyRequest struct{}"
	}

	return strings.Join([]string{"ListDomainPermissionsForAgencyRequest", string(data)}, " ")
}
