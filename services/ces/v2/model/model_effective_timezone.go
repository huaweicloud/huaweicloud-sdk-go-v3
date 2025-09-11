package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EffectiveTimezone **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **约束限制**： 不涉及。 **取值范围**： 长度为[1,16]个字符。           **默认取值**： 不涉及。
type EffectiveTimezone struct {
}

func (o EffectiveTimezone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EffectiveTimezone struct{}"
	}

	return strings.Join([]string{"EffectiveTimezone", string(data)}, " ")
}
