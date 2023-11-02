package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResizeOrderRequestBody 包周期实例变更规格请求
type CreateResizeOrderRequestBody struct {

	// 自动开启安全组规则
	AutoOpenSecurityGroupRule *bool `json:"auto_open_security_group_rule,omitempty"`

	// 启动迅速
	ExecuteImmediately *bool `json:"execute_immediately,omitempty"`

	// 新的容量，单位是GB
	NewCapacity *float32 `json:"new_capacity,omitempty"`

	// 密码
	Password *string `json:"password,omitempty"`

	// 区域代码
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o CreateResizeOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResizeOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResizeOrderRequestBody", string(data)}, " ")
}
