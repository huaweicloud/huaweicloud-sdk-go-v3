package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EffectiveTimezoneResp **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **取值范围**： 长度为[1,16]个字符。
type EffectiveTimezoneResp struct {
}

func (o EffectiveTimezoneResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EffectiveTimezoneResp struct{}"
	}

	return strings.Join([]string{"EffectiveTimezoneResp", string(data)}, " ")
}
