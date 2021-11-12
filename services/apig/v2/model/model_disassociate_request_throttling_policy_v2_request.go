package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateRequestThrottlingPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API和流控策略绑定关系的ID。

	ThrottleBindingId string `json:"throttle_binding_id"`
}

func (o DisassociateRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
