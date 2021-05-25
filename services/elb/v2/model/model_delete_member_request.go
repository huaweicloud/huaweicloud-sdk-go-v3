package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMemberRequest struct {
	// 后端云服务器组id

	PoolId string `json:"pool_id"`
	// 后端云服务器id

	MemberId string `json:"member_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
