package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteScaleOutPolicyRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o DeleteScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteScaleOutPolicyRequest", string(data)}, " ")
}
