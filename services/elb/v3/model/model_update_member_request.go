package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMemberRequest struct {
	// 后端服务器ID。

	MemberId string `json:"member_id"`
	// 后端服务器组ID。

	PoolId string `json:"pool_id"`

	Body *UpdateMemberRequestBody `json:"body,omitempty"`
}

func (o UpdateMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequest", string(data)}, " ")
}
