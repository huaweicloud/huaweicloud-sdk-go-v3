package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionInfo struct {

	// 动作类型：block:阻断,captcha:人机验证;log: 仅记录;dynamic_block:动态阻断
	Category *string `json:"category,omitempty"`

	Detail *DetailInfo `json:"detail,omitempty"`
}

func (o ActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionInfo struct{}"
	}

	return strings.Join([]string{"ActionInfo", string(data)}, " ")
}
