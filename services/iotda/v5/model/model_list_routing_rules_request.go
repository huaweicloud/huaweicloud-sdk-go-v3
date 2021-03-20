package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRoutingRulesRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Resource *string `json:"resource,omitempty"`

	Event *string `json:"event,omitempty"`

	AppType *string `json:"app_type,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	RuleName *string `json:"rule_name,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRoutingRulesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRoutingRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRoutingRulesRequest", string(data)}, " ")
}
