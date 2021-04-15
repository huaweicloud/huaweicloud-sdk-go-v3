package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronListSecurityGroupRulesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *string `json:"id,omitempty"`

	Direction *string `json:"direction,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	Ethertype *string `json:"ethertype,omitempty"`

	Description *string `json:"description,omitempty"`

	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`

	RemoteGroupId *string `json:"remote_group_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	PortRangeMax *string `json:"port_range_max,omitempty"`

	PortRangeMin *string `json:"port_range_min,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NeutronListSecurityGroupRulesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupRulesRequest struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupRulesRequest", string(data)}, " ")
}
