package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotEffectHostNum **参数解释**: 学习完成策略未生效主机数 **取值范围**: 最小值0，最大值2147483647
type NotEffectHostNum struct {
}

func (o NotEffectHostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotEffectHostNum struct{}"
	}

	return strings.Join([]string{"NotEffectHostNum", string(data)}, " ")
}
