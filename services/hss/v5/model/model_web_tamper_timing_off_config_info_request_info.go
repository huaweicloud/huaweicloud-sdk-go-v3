package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperTimingOffConfigInfoRequestInfo 定时关闭防护功能配置详情
type WebTamperTimingOffConfigInfoRequestInfo struct {

	// **参数解释**: 自动关闭防护周期列表 **约束限制**: 不涉及 **取值范围**: 最少1条，最多7条 **默认取值**: 不涉及
	WeekOffList *[]int32 `json:"week_off_list,omitempty"`

	// **参数解释**: 自动关闭防护时间段 **约束限制**: 不涉及 **取值范围**: 最少1条，最多5条 **默认取值**: 不涉及
	TimingRangeList *[]TimingRangeConfigRequestInfo `json:"timing_range_list,omitempty"`
}

func (o WebTamperTimingOffConfigInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperTimingOffConfigInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperTimingOffConfigInfoRequestInfo", string(data)}, " ")
}
