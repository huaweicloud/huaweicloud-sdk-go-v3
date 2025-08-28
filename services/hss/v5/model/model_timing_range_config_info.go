package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimingRangeConfigInfo struct {

	// **参数解释**: 自动关闭防护的时间段范围 **取值范围**: 字符长度0-512位
	TimeRange *string `json:"time_range,omitempty"`

	// **参数解释**: 自动关闭防护的时间段描述 **取值范围**: 字符长度0-512位
	Description *string `json:"description,omitempty"`
}

func (o TimingRangeConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimingRangeConfigInfo struct{}"
	}

	return strings.Join([]string{"TimingRangeConfigInfo", string(data)}, " ")
}
