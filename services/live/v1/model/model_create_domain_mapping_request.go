package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDomainMappingRequest struct {
	SpecifyProject *string `json:"specify_project,omitempty"`

	Body *DomainMapping `json:"body,omitempty"`
}

func (o CreateDomainMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDomainMappingRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainMappingRequest", string(data)}, " ")
}
