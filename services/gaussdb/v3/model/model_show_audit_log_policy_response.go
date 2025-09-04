package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditLogPolicyResponse Response Object
type ShowAuditLogPolicyResponse struct {

	// **参数解释**：  审计日志保存天数，0表示关闭审计日志策略。   **取值范围**：  0~732。
	KeepDays *int32 `json:"keep_days,omitempty"`

	// **参数解释**：  审计记录的操作类型。空表示不过滤任何操作类型。  **取值范围**：  动态范围。
	AuditTypes *[]string `json:"audit_types,omitempty"`

	// **参数解释**：  审计记录的操作类型。空表示不过滤任何操作类型。  **取值范围**：  不涉及
	AllAuditLogAction *string `json:"all_audit_log_action,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ShowAuditLogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditLogPolicyResponse", string(data)}, " ")
}
