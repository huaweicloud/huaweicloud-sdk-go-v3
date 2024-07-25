package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsPvo struct {

	// **参数解释：**  结束时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EndTime string `json:"endTime"`

	// **参数解释：**  开始时间。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	StartTime string `json:"startTime"`
}

func (o StatisticsPvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsPvo struct{}"
	}

	return strings.Join([]string{"StatisticsPvo", string(data)}, " ")
}
