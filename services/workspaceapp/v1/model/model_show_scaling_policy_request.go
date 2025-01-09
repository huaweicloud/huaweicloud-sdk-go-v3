package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScalingPolicyRequest Request Object
type ShowScalingPolicyRequest struct {

	// 服务器组唯一标识。
	ServerGroupId string `json:"server_group_id"`
}

func (o ShowScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingPolicyRequest", string(data)}, " ")
}
