package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntelligentKillSessionStatistic 实时统计信息
type IntelligentKillSessionStatistic struct {

	// **参数解释**：  该类统计维度下提取到的限流关键字。  **约束限制**：  不涉及。
	Keyword *string `json:"keyword,omitempty"`

	// **参数解释**：  随机选取符合sql限流关键字的用户某条sql样例。  **约束限制**：  不涉及。
	RawSql *string `json:"raw_sql,omitempty"`

	// **参数解释**：  符合该统计维度的线程id。  **约束限制**：  不涉及。
	Ids *[]int64 `json:"ids,omitempty"`

	// **参数解释**：  符合该统计维度的线程id总数量。  **约束限制**：  不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  符合该统计维度的线程总执行时间。  **约束限制**：  不涉及。
	TotalTime *float64 `json:"total_time,omitempty"`

	// **参数解释**：  符合该统计维度的线程平均执行时间。  **约束限制**：  不涉及。
	AvgTime *float64 `json:"avg_time,omitempty"`

	// **参数解释**：  符合该统计维度的线程最大执行时间。  **约束限制**：  不涉及。
	MaxTime *float64 `json:"max_time,omitempty"`

	// **参数解释**：  统计维度。  **约束限制**：  不涉及。
	Strategy *string `json:"strategy,omitempty"`

	// **参数解释**：  推荐最大并发数。type为\"kill\"时该参数没有返回值。  **约束限制**：  不涉及。
	AdviceConcurrency *int32 `json:"advice_concurrency,omitempty"`

	// **参数解释**：  该条维度数据的类型。\"kill\"表示当前统计时刻下一键kill会话下发后会kill该类会话；\"limit\"表示当前统计时刻下勾选\"同步开启sql限流和添加规则\"时会添加的规则。  **约束限制**：  不涉及。
	Type *string `json:"type,omitempty"`
}

func (o IntelligentKillSessionStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntelligentKillSessionStatistic struct{}"
	}

	return strings.Join([]string{"IntelligentKillSessionStatistic", string(data)}, " ")
}
