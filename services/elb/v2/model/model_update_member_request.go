package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMemberRequest struct {

	// 后端云服务器id
	MemberId string `json:"member_id" xml:"member_id"`

	// 后端云服务器组id
	PoolId string `json:"pool_id" xml:"pool_id"`

	Body *UpdateMemberRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequest", string(data)}, " ")
}
