package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoricalSqlFilterRuleRequest Request Object
type ShowHistoricalSqlFilterRuleRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  节点ID，此参数是节点的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。  **默认取值**：  不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**：  SQL限流类型。  **约束限制**：  不涉及。  **取值范围**：  - SELECT：查询语句。 - UPDATE：更新语句。 - DELETE：删除语句。 - INSERT：插入语句。  **默认取值**：  不传则默认查询所有类型的限流规则。
	SqlType *string `json:"sql_type,omitempty"`
}

func (o ShowHistoricalSqlFilterRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoricalSqlFilterRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoricalSqlFilterRuleRequest", string(data)}, " ")
}
