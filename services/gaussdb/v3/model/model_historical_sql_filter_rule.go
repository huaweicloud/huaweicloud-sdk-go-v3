package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistoricalSqlFilterRule **参数解释**：  SQL限流规则。
type HistoricalSqlFilterRule struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID。  **取值范围**：  与入参中的节点ID一致。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  SQL限流规则。  **取值范围**：  由一个或多个关键字（最多为128个关键字）组成，关键字之间通过\"~\"分隔符分开，如select~from~t1。规则中不能包含‘\\’、中英文逗号、‘~~’，不能以‘~’结尾。  **默认取值**：  不涉及
	Pattern string `json:"pattern"`

	// **参数解释**：  SQL限流类型。  **取值范围**：  - SELECT：查询语句。 - UPDATE：更新语句。 - DELETE：删除语句。 - INSERT：插入语句。  **默认取值**：  不涉及。
	SqlType string `json:"sql_type"`

	// **参数解释**：  最大并发数。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	MaxConcurrency int32 `json:"max_concurrency"`

	// **参数解释**：  SQL限流规则创建时间。  **取值范围**：  0 - 9223372036854775807。  **默认取值**：  不涉及
	CreateAt int64 `json:"create_at"`

	// **参数解释**：  SQL限流规则失效时间。  **取值范围**：  0 - 9223372036854775807。  **默认取值**：  不涉及
	ExpireAt int64 `json:"expire_at"`

	// **参数解释**：  SQL限流规则删除时间。  **取值范围**：  0 - 9223372036854775807。  **默认取值**：  不涉及
	DeleteAt int64 `json:"delete_at"`
}

func (o HistoricalSqlFilterRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoricalSqlFilterRule struct{}"
	}

	return strings.Join([]string{"HistoricalSqlFilterRule", string(data)}, " ")
}
