package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlFilterRulePattern SQL限流规则和最大并发数
type SqlFilterRulePattern struct {

	// **参数解释**：  SQL限流规则。  **取值范围**：  由一个或多个关键字（最多为128个关键字）组成，关键字之间通过\"~\"分隔符分开，如select~from~t1。规则中不能包含‘\\’、中英文逗号、‘~~’，不能以‘~’结尾。
	Pattern string `json:"pattern"`

	// **参数解释**：  最大并发数。  **取值范围**：  不涉及。
	MaxConcurrency int32 `json:"max_concurrency"`

	// **参数解释**：  当前并发数。  **取值范围**：  0 - 4294967296。
	CurConcurrency *int32 `json:"cur_concurrency,omitempty"`

	// **参数解释**：  当前拦截次数。  **取值范围**：  0 - 4294967296。
	CurReject *int32 `json:"cur_reject,omitempty"`

	// **参数解释**：  SQL限流规则创建时间。  **取值范围**：  0 - 9223372036854775807。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：  SQL限流规则失效时间。  **取值范围**：  0 - 9223372036854775807。
	ExpireAt *int64 `json:"expire_at,omitempty"`
}

func (o SqlFilterRulePattern) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlFilterRulePattern struct{}"
	}

	return strings.Join([]string{"SqlFilterRulePattern", string(data)}, " ")
}
