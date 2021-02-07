package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDomainMappingRequest struct {
	SpecifyProject *string `json:"specify_project,omitempty"`
	PullDomain     string  `json:"pull_domain"`
	PushDomain     string  `json:"push_domain"`
}

func (o DeleteDomainMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDomainMappingRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainMappingRequest", string(data)}, " ")
}
