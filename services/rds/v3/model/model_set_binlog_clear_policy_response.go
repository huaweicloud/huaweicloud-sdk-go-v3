package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SetBinlogClearPolicyResponse struct {
	// 操作结果。

	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetBinlogClearPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetBinlogClearPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetBinlogClearPolicyResponse", string(data)}, " ")
}
