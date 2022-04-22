package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleBindingBatchDelete struct {

	// 需要解除绑定的API和流控策略绑定关系ID列表
	ThrottleBindings *[]string `json:"throttle_bindings,omitempty"`
}

func (o ThrottleBindingBatchDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleBindingBatchDelete struct{}"
	}

	return strings.Join([]string{"ThrottleBindingBatchDelete", string(data)}, " ")
}
