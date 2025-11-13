package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSqlFilterRulePattern SQL限流规则和最大并发数
type NodeSqlFilterRulePattern struct {

	// SQL限流规则，由一个或多个关键字（最多为128个关键字）组成，关键字之间通过\"~\"分隔符分开，如select~from~t1。规则中不能包含‘\\’、中英文逗号、‘~~’，不能以‘~’结尾。
	Pattern string `json:"pattern"`

	// 最大并发数。取值范围：非负整数。
	MaxConcurrency int32 `json:"max_concurrency"`

	// **参数解释**：  SQL限流失效时间，标准秒级时间戳，永久生效SQL限流规则该字段为null。  **约束限制**：  不涉及。  **取值范围**：  0 - 9223372036854775807。  **默认取值**：  不涉及。
	ExpireAt *int64 `json:"expire_at,omitempty"`
}

func (o NodeSqlFilterRulePattern) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSqlFilterRulePattern struct{}"
	}

	return strings.Join([]string{"NodeSqlFilterRulePattern", string(data)}, " ")
}
