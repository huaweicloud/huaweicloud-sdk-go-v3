package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartRecyclePolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *RecyclePolicyRequestBody `json:"body,omitempty"`
}

func (o StartRecyclePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"StartRecyclePolicyRequest", string(data)}, " ")
}
