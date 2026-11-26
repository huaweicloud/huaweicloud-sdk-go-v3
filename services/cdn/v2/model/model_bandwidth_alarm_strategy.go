package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthAlarmStrategy **参数解释：** 带宽告警策略，当带宽达到告警阈值时，发送告警通知 **约束限制：** 不涉及
type BandwidthAlarmStrategy struct {

	// **参数解释：** 带宽告警开关 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** off: 关闭
	Status *string `json:"status,omitempty"`

	// **参数解释：** 带宽告警阈值，单位：bit/s **约束限制：** - status为on时必填  **取值范围：** 1000000-10000000000000 **默认取值：** 不涉及
	Value *int64 `json:"value,omitempty"`
}

func (o BandwidthAlarmStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthAlarmStrategy struct{}"
	}

	return strings.Join([]string{"BandwidthAlarmStrategy", string(data)}, " ")
}
