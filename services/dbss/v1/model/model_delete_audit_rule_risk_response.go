package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditRuleRiskResponse Response Object
type DeleteAuditRuleRiskResponse struct {

	// 状态  - SUCCESS：成功  - FAILED：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditRuleRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditRuleRiskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditRuleRiskResponse", string(data)}, " ")
}
