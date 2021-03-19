package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAllMembersRequest struct {
	Address *[]string `json:"address,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	IpVersion *string `json:"ip_version,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Name *[]string `json:"name,omitempty"`

	OperatingStatus *[]string `json:"operating_status,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	PoolId *string `json:"pool_id,omitempty"`

	ProtocolPort *[]int32 `json:"protocol_port,omitempty"`

	SubnetCidrId *[]string `json:"subnet_cidr_id,omitempty"`

	Weight *[]int32 `json:"weight,omitempty"`
}

func (o ListAllMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllMembersRequest struct{}"
	}

	return strings.Join([]string{"ListAllMembersRequest", string(data)}, " ")
}
