package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScaleOutPolicyRequest Request Object
type UpdateScaleOutPolicyRequest struct {

	// 策略id
	Id string `json:"id"`

	Body *UpdateScaleOutPolicyReq `json:"body,omitempty"`
}

func (o UpdateScaleOutPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScaleOutPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScaleOutPolicyRequest", string(data)}, " ")
}
