package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMemberRequest struct {
	// 后端服务器ID。

	MemberId string `json:"member_id"`
	// 后端服务器组ID。

	PoolId string `json:"pool_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
