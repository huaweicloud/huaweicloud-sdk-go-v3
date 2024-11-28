package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisPitrPolicyRequest Request Object
type ShowRedisPitrPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowRedisPitrPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisPitrPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowRedisPitrPolicyRequest", string(data)}, " ")
}
