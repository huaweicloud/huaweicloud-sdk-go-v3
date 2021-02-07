package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMembersRequest struct {
	PoolId       string  `json:"pool_id"`
	Limit        *int32  `json:"limit,omitempty"`
	Marker       *string `json:"marker,omitempty"`
	PageReverse  *bool   `json:"page_reverse,omitempty"`
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Address      *string `json:"address,omitempty"`
	ProtocolPort *int32  `json:"protocol_port,omitempty"`
	SubnetId     *string `json:"subnet_id,omitempty"`
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
	Weight       *int32  `json:"weight,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
