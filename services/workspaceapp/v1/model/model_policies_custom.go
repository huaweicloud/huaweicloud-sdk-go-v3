package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesCustom 自定义策略
type PoliciesCustom struct {

	// 自定义策略配置项1： false: 表示关闭 true: 表示开启
	CustomConfiguration1Enable *bool `json:"custom_configuration1_enable,omitempty"`

	Options *PoliciesCustomOptions `json:"options,omitempty"`
}

func (o PoliciesCustom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesCustom struct{}"
	}

	return strings.Join([]string{"PoliciesCustom", string(data)}, " ")
}
