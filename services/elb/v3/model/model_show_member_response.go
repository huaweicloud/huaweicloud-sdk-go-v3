package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMemberResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberResponse", string(data)}, " ")
}
