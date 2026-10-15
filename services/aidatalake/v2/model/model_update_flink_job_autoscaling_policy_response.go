package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkJobAutoscalingPolicyResponse Response Object
type UpdateFlinkJobAutoscalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateFlinkJobAutoscalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJobAutoscalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJobAutoscalingPolicyResponse", string(data)}, " ")
}
