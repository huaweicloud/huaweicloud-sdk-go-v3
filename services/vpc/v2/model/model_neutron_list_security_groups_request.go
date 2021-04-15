package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronListSecurityGroupsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NeutronListSecurityGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupsRequest struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupsRequest", string(data)}, " ")
}
