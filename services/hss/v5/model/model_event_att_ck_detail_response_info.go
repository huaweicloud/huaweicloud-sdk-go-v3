package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventAttCkDetailResponseInfo 各种ATT&CK攻击阶段的名称与数量列表
type EventAttCkDetailResponseInfo struct {

	// **参数解释**： 攻击阶段名称，根据页面语言环境，返回中文或英文 **取值范围**： 字符长度1-64位
	AttCk *string `json:"att_ck,omitempty"`

	// **参数解释**: 数量 **取值范围**: 最小值0，最大值2147483647
	Num *int32 `json:"num,omitempty"`
}

func (o EventAttCkDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventAttCkDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"EventAttCkDetailResponseInfo", string(data)}, " ")
}
