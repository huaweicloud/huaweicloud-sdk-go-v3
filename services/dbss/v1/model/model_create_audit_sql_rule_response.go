package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuditSqlRuleResponse Response Object
type CreateAuditSqlRuleResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuditSqlRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuditSqlRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAuditSqlRuleResponse", string(data)}, " ")
}
