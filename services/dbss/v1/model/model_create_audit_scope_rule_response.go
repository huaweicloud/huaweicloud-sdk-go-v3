package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuditScopeRuleResponse Response Object
type CreateAuditScopeRuleResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuditScopeRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuditScopeRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAuditScopeRuleResponse", string(data)}, " ")
}
