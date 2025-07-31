package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessStartTime **参数解释**: 进程启动时间 **取值范围**: 最小值0，最大值9223372036854775807
type ProcessStartTime struct {
}

func (o ProcessStartTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessStartTime struct{}"
	}

	return strings.Join([]string{"ProcessStartTime", string(data)}, " ")
}
