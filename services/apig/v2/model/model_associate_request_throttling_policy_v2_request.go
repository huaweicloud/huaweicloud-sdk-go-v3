package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRequestThrottlingPolicyV2Request Request Object
type AssociateRequestThrottlingPolicyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ThrottleApiBindingCreate `json:"body,omitempty"`
}

func (o AssociateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
