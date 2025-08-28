package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimingRangeConfigRequestInfo struct {

	// **参数解释**: 自动关闭防护的时间段范围 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
	TimeRange *string `json:"time_range,omitempty"`

	// **参数解释**: 自动关闭防护的时间段描述 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
	Description *string `json:"description,omitempty"`
}

func (o TimingRangeConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimingRangeConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"TimingRangeConfigRequestInfo", string(data)}, " ")
}
