package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAllProjectsPermissionsForAgencyRequest struct {
	AgencyId string `json:"agency_id"`
	DomainId string `json:"domain_id"`
}

func (o ListAllProjectsPermissionsForAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllProjectsPermissionsForAgencyRequest struct{}"
	}

	return strings.Join([]string{"ListAllProjectsPermissionsForAgencyRequest", string(data)}, " ")
}
