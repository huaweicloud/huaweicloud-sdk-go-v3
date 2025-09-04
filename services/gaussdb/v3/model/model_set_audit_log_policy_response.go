package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditLogPolicyResponse Response Object
type SetAuditLogPolicyResponse struct {

	// **参数解释**：  设置审计日志策略的操作结果。  **取值范围**：  COMPLETED|FAILED。
	Result *string `json:"result,omitempty"`

	// **参数解释**：  任务流id。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditLogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditLogPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAuditLogPolicyResponse", string(data)}, " ")
}
