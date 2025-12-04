package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyEipRequest Request Object
type ModifyEipRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyElbVipOpenReq `json:"body,omitempty"`
}

func (o ModifyEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEipRequest struct{}"
	}

	return strings.Join([]string{"ModifyEipRequest", string(data)}, " ")
}
