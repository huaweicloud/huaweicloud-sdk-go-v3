package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditSqlRuleResponse Response Object
type UpdateAuditSqlRuleResponse struct {

	// 状态  - SUCCESS：成功  - FAILED：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditSqlRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditSqlRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditSqlRuleResponse", string(data)}, " ")
}
