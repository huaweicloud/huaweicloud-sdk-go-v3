package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EffectiveTimezone 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"
type EffectiveTimezone struct {
}

func (o EffectiveTimezone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EffectiveTimezone struct{}"
	}

	return strings.Join([]string{"EffectiveTimezone", string(data)}, " ")
}
