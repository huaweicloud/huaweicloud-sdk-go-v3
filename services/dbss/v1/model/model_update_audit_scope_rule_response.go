package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditScopeRuleResponse Response Object
type UpdateAuditScopeRuleResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditScopeRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditScopeRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditScopeRuleResponse", string(data)}, " ")
}
