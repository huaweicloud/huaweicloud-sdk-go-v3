package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRulesRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	AppId *string `json:"app_id,omitempty"`

	RuleType *string `json:"rule_type,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRulesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRulesRequest", string(data)}, " ")
}
