package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScalingPolicyRequest Request Object
type DeleteScalingPolicyRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingPolicyRequest", string(data)}, " ")
}
