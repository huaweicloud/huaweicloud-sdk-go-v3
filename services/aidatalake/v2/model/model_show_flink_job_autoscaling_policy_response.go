package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkJobAutoscalingPolicyResponse Response Object
type ShowFlinkJobAutoscalingPolicyResponse struct {
	AutoscalingPolicy *AutoscalingPolicyResponse `json:"autoscaling_policy,omitempty"`
	HttpStatusCode    int                        `json:"-"`
}

func (o ShowFlinkJobAutoscalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobAutoscalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobAutoscalingPolicyResponse", string(data)}, " ")
}
