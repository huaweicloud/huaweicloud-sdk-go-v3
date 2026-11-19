package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmTotalNumRes **参数解释**: 告警总数 **取值范围**: 最小值0，最大值2147483647
type AlarmTotalNumRes struct {
}

func (o AlarmTotalNumRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTotalNumRes struct{}"
	}

	return strings.Join([]string{"AlarmTotalNumRes", string(data)}, " ")
}
