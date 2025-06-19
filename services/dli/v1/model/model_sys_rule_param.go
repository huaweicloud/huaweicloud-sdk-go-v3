package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SysRuleParam struct {

	// 阀值上限
	Max *int64 `json:"max,omitempty"`

	// 阀值下限
	Min *int64 `json:"min,omitempty"`

	// 阀值默认值
	DefaultValue *int64 `json:"defaultValue,omitempty"`

	// 描述
	Desc *string `json:"desc,omitempty"`
}

func (o SysRuleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysRuleParam struct{}"
	}

	return strings.Join([]string{"SysRuleParam", string(data)}, " ")
}
