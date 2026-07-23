package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsPvo struct {

	// **参数解释：**  统计区间的开始时间，用于指定统计时间区间的起始点。  **约束限制：**  不能为空，且必须早于或等于endTime。  **取值范围：**  UTC标准时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **默认取值：**  不涉及。
	StartTime string `json:"startTime"`

	// **参数解释：**  统计区间的结束时间，用于指定统计时间区间的结束点。  **约束限制：**  不能为空，且必须晚于或等于startTime。  **取值范围：**  UTC标准时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **默认取值：**  不涉及。
	EndTime string `json:"endTime"`
}

func (o StatisticsPvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsPvo struct{}"
	}

	return strings.Join([]string{"StatisticsPvo", string(data)}, " ")
}
