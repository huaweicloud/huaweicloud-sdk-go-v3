package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuditRiskRuleResponse Response Object
type CreateAuditRiskRuleResponse struct {

	// 状态  - SUCCESS：成功  - FAILED：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuditRiskRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuditRiskRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAuditRiskRuleResponse", string(data)}, " ")
}
