package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRecordRulesResponse struct {
	// 录制规则列表

	Rules *[]RecordRule `json:"rules,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecordRulesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRecordRulesResponse", string(data)}, " ")
}
