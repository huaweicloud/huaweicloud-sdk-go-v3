package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateThrottlingPolicyV2Request Request Object
type BatchDisassociateThrottlingPolicyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 必须为delete
	Action string `json:"action"`

	Body *ThrottleBindingBatchDelete `json:"body,omitempty"`
}

func (o BatchDisassociateThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"BatchDisassociateThrottlingPolicyV2Request", string(data)}, " ")
}
