package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesCustom 自定义策略。
type PoliciesCustom struct {

	// 自定义配置启用。
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
