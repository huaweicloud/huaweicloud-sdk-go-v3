package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopScaleOutPolicyRequest Request Object
type StopScaleOutPolicyRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o StopScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"StopScaleOutPolicyRequest", string(data)}, " ")
}
