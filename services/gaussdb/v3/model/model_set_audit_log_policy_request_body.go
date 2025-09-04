package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetAuditLogPolicyRequestBody struct {

	// **参数解释**：  审计日志保存天数，0表示关闭审计日志策略。  **约束限制**：  不涉及。  **取值范围**：  0~732。  **默认取值**：  7。
	KeepDays int32 `json:"keep_days"`

	// **参数解释**：  仅关闭审计日志策略时有效。 - true（默认），表示关闭审计日志策略的同时，保留历史审计日志。 - false，表示关闭审计日志策略的同时，删除已有的历史审计日志。  **约束限制**：  不涉及。  **取值范围**：  true|false。  **默认取值**：  true。
	ReserveAuditLogs *bool `json:"reserve_audit_logs,omitempty"`

	// **参数解释**：  审计记录的操作类型，动态范围。空表示不过滤任何操作类型。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	AuditTypes *[]string `json:"audit_types,omitempty"`
}

func (o SetAuditLogPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditLogPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetAuditLogPolicyRequestBody", string(data)}, " ")
}
