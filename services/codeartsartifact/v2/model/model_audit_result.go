package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditResult struct {

	// **参数解释**： 审计日志列表。  **取值范围**： 不涉及。
	AuditInfoList *[]AuditDo `json:"auditInfoList,omitempty"`

	// **参数解释**： 审计日志条数。  **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o AuditResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditResult struct{}"
	}

	return strings.Join([]string{"AuditResult", string(data)}, " ")
}
