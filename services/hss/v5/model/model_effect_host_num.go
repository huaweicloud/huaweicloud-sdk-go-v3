package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EffectHostNum **参数解释**: 学习完成策略已生效主机数 **取值范围**: 最小值0，最大值2147483647
type EffectHostNum struct {
}

func (o EffectHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EffectHostNum struct{}"
	}

	return strings.Join([]string{"EffectHostNum", string(data)}, " ")
}
