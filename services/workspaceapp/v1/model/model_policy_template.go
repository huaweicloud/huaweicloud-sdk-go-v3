package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyTemplate 策略模板
type PolicyTemplate struct {

	// 策略组名称，名称需要满足如下规则 1. 不能有空格和.开头，可以由中文、数字、英文大小写组成 2. 字符长度范围1-55
	PolicyGroupName string `json:"policy_group_name"`

	// 描述
	Description *string `json:"description,omitempty"`

	Policies *Policies `json:"policies"`
}

func (o PolicyTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTemplate struct{}"
	}

	return strings.Join([]string{"PolicyTemplate", string(data)}, " ")
}
