package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogDesensitizationRequest Request Object
type ShowSlowLogDesensitizationRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowSlowLogDesensitizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDesensitizationRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDesensitizationRequest", string(data)}, " ")
}
