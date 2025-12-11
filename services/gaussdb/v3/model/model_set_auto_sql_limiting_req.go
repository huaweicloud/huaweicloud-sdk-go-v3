package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoSqlLimitingReq **参数解释**：  开启自治限流规则请求体。  **约束限制**：  不涉及。
type SetAutoSqlLimitingReq struct {

	// **参数解释**：  开启自治限流规则。  **约束限制**：  不涉及。
	AutoSqlLimitingRules []AutoSqlLimitingRule `json:"auto_sql_limiting_rules"`
}

func (o SetAutoSqlLimitingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoSqlLimitingReq struct{}"
	}

	return strings.Join([]string{"SetAutoSqlLimitingReq", string(data)}, " ")
}
