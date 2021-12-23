package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRequestThrottlingPolicyV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *ThrottleBaseInfo `json:"body,omitempty"`
}

func (o CreateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"CreateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
