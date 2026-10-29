package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogCounts 参数集
type LogCounts struct {

	// **参数解释：** 子时间区间的日志条数。 **取值范围：** 不涉及。
	Num *int32 `json:"num,omitempty"`

	// **参数解释：** 子时间区间的起始时间。 **取值范围：** 不涉及。
	StartTime *int64 `json:"startTime,omitempty"`

	// **参数解释：** 子时间区间的结束时间。 **取值范围：** 不涉及。
	EndTime *int64 `json:"endTime,omitempty"`
}

func (o LogCounts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogCounts struct{}"
	}

	return strings.Join([]string{"LogCounts", string(data)}, " ")
}
