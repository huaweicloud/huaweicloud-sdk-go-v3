package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRuleActionsRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	RuleId *string `json:"rule_id,omitempty"`

	Channel *string `json:"channel,omitempty"`

	AppType *string `json:"app_type,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRuleActionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRuleActionsRequest struct{}"
	}

	return strings.Join([]string{"ListRuleActionsRequest", string(data)}, " ")
}
