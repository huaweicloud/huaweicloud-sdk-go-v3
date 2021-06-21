package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRecordRuleRequest struct {
	Body *RecordRuleRequest `json:"body,omitempty"`
}

func (o CreateRecordRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordRuleRequest", string(data)}, " ")
}
