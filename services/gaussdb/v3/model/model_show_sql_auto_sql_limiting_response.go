package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlAutoSqlLimitingResponse Response Object
type ShowSqlAutoSqlLimitingResponse struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**  节点的自治限流规则。
	AutoSqlLimitingRules *[]NodeSqlAutoSqlLimiting `json:"auto_sql_limiting_rules,omitempty"`
	HttpStatusCode       int                       `json:"-"`
}

func (o ShowSqlAutoSqlLimitingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlAutoSqlLimitingResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlAutoSqlLimitingResponse", string(data)}, " ")
}
