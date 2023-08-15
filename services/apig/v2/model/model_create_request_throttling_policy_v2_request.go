package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRequestThrottlingPolicyV2Request Request Object
type CreateRequestThrottlingPolicyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
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
