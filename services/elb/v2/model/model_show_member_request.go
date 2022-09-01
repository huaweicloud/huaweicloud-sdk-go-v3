package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMemberRequest struct {

	// 后端云服务器组id
	PoolId string `json:"pool_id" xml:"pool_id"`

	// 后端云服务器id
	MemberId string `json:"member_id" xml:"member_id"`
}

func (o ShowMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberRequest", string(data)}, " ")
}
