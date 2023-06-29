package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoScalingPolicyResponse Response Object
type ShowAutoScalingPolicyResponse struct {

	// 弹性伸缩策略列表
	Body           *[]AutoScalingPolicyV2 `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoScalingPolicyResponse", string(data)}, " ")
}
