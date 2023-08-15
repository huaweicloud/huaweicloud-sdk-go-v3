package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateRequestThrottlingPolicyV2Request Request Object
type DisassociateRequestThrottlingPolicyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// API和流控策略绑定关系的编号
	ThrottleBindingId string `json:"throttle_binding_id"`
}

func (o DisassociateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
