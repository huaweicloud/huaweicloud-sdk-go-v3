package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMemberRequest Request Object
type DeleteMemberRequest struct {

	// 后端服务器组ID。
	PoolId string `json:"pool_id"`

	// 后端服务器ID。 >说明： 此处并非ECS服务器的ID，而是ELB为绑定的后端服务器自动生成的member ID。
	MemberId string `json:"member_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
