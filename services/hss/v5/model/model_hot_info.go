package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HotInfo struct {

	// **参数解释**: 标题 **取值范围**: 字符长度0-256位
	Title *string `json:"title,omitempty"`

	// **参数解释**: 更新时间 **取值范围**: 最小值0，最大值4071095999000
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 级别 **取值范围**: 字符长度0-32位
	SeverityLevel *string `json:"severity_level,omitempty"`
}

func (o HotInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotInfo struct{}"
	}

	return strings.Join([]string{"HotInfo", string(data)}, " ")
}
