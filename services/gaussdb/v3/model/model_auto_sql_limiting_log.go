package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoSqlLimitingLog **参数解释**：  自治限流执行记录。  **约束限制**：  不涉及
type AutoSqlLimitingLog struct {

	// **参数解释**：  节点ID。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**：  SQL限流规则。  **取值范围**：  由一个或多个关键字（最多为128个关键字）组成，关键字之间通过\"~\"分隔符分开，如select~from~t1。规则中不能包含‘\\’、中英文逗号、‘~~’，不能以‘~’结尾。  **默认取值**：  不涉及。
	Pattern *string `json:"pattern,omitempty"`

	// **参数解释**：  限流类型。  **取值范围**：  SELECT。  **默认取值**：  不涉及。
	SqlType *string `json:"sql_type,omitempty"`

	// **参数解释**：  最大并发数。  **取值范围**：  [0~1000000000]。  **默认取值**：  不涉及。
	MaxConcurrency *int32 `json:"max_concurrency,omitempty"`

	// **参数解释**：  限流开始时间戳。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：  限流失效时间戳。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ExpireAt *int64 `json:"expire_at,omitempty"`
}

func (o AutoSqlLimitingLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoSqlLimitingLog struct{}"
	}

	return strings.Join([]string{"AutoSqlLimitingLog", string(data)}, " ")
}
