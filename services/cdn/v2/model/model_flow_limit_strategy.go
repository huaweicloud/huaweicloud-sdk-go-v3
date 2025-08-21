package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowLimitStrategy **参数解释：** 用量封顶配置 **约束限制：** 不涉及
type FlowLimitStrategy struct {

	// **参数解释：**  用量统计类型 **约束限制：** 不涉及 **取值范围：** - instant: 瞬时用量 - hour: 累计用量（小时） - day: 累计用量（天） **默认取值：** 不涉及
	StrategyType *string `json:"strategy_type,omitempty"`

	// **参数解释：**  用量封顶类型 **约束限制：** 不涉及 **取值范围：** - bandwidth: 带宽封顶，单位: bit/s - traffic: 流量封顶，单位: bit **默认取值：** 不涉及
	ItemType *string `json:"item_type,omitempty"`

	// **参数解释：** 用量封顶阈值，域名用量达到该阈值后，将会停用域名 **约束限制：** 不涉及 **取值范围：** 必须为正整数 **默认取值：** 不涉及
	LimitValue *string `json:"limit_value,omitempty"`

	// **参数解释：** 用量告警阈值，域名用量达到该阈值后，将会发送告警 **约束限制：** 不涉及 **取值范围：** 1-90 **默认取值：** 不涉及
	AlarmPercentThreshold *string `json:"alarm_percent_threshold,omitempty"`

	// **参数解释：** 域名封禁周期 **约束限制：** ban_time设置为0时，表示不自动解封，需要客户手动解封域名 **取值范围：** - 0: 不自动解封 - 60: 60分钟，即1个小时 - 720: 720分钟，即12个小时 - 1440: 1440分钟，即24个小时 - 4320: 4320分钟，即3天 **默认取值：** 不涉及
	BanTime *string `json:"ban_time,omitempty"`
}

func (o FlowLimitStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowLimitStrategy struct{}"
	}

	return strings.Join([]string{"FlowLimitStrategy", string(data)}, " ")
}
