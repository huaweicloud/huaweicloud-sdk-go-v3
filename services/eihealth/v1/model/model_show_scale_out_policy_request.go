package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScaleOutPolicyRequest Request Object
type ShowScaleOutPolicyRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o ShowScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScaleOutPolicyRequest", string(data)}, " ")
}
