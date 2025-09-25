package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlTypeRangeConfigResult 全量SQL开关记录
type SqlTypeRangeConfigResult struct {

	// **参数解释**: 是否开启全量SQL。 **取值范围**: - true：已开启。 - false：已关闭。
	IsOpen *bool `json:"is_open,omitempty"`

	// **参数解释**: 开关状态持续的开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。 **取值范围**: 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`
}

func (o SqlTypeRangeConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlTypeRangeConfigResult struct{}"
	}

	return strings.Join([]string{"SqlTypeRangeConfigResult", string(data)}, " ")
}
