package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMemberRequest struct {
	// 后端服务器ID。

	MemberId string `json:"member_id"`
	// 后端服务器组ID。

	PoolId string `json:"pool_id"`
}

func (o ShowMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberRequest", string(data)}, " ")
}
