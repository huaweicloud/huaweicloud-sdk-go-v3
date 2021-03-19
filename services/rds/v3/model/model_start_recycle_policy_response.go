package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartRecyclePolicyResponse struct {
	// 操作结果。

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartRecyclePolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"StartRecyclePolicyResponse", string(data)}, " ")
}
