package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequestLimitRulesEngine **参数解释：** 请求限速，将用户请求速度限制在指定范围内，一定程度上减少突发高带宽风险，节省成本 **约束限制：** 不涉及
type RequestLimitRulesEngine struct {

	// **参数解释：** 限速条件 > 例如：type=size，limit_rate_after=50表示从传输50个字节后开始限速且限速值为limit_rate_value  **约束限制：** 不涉及 **取值范围：** 0-1073741824，单位：byte **默认取值：** 不涉及
	LimitRateAfter int64 `json:"limit_rate_after"`

	// **参数解释：** 限速值，即达到限速条件后的最大访问速度 **约束限制：** 不涉及 **取值范围：** 0-104857600，单位：Bps **默认取值：** 不涉及
	LimitRateValue int32 `json:"limit_rate_value"`
}

func (o RequestLimitRulesEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestLimitRulesEngine struct{}"
	}

	return strings.Join([]string{"RequestLimitRulesEngine", string(data)}, " ")
}
