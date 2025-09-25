package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimingRangeConfigRequestInfo 自动关闭防护时间段
type TimingRangeConfigRequestInfo struct {

	// **参数解释**: 自动关闭防护的时间段范围，开始时间与结束之间用中划线相连，如15:00-15:30。 **约束限制**: 每个时间段最短不能少于5分钟。多个时间段之间不允许重叠，且两段时间间隔必须大于5分钟（时间00:00和23:59特例除外）。 **取值范围**: 字符长度0-20位 **默认取值**: 不涉及
	TimeRange *string `json:"time_range,omitempty"`

	// **参数解释**: 自动关闭防护的时间段描述 **约束限制**: 不涉及 **取值范围**: 字符长度0-20位 **默认取值**: 不涉及
	Description *string `json:"description,omitempty"`
}

func (o TimingRangeConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimingRangeConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"TimingRangeConfigRequestInfo", string(data)}, " ")
}
