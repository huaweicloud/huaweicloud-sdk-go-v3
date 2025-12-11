package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditScopeRuleSwitchResponse Response Object
type SetAuditScopeRuleSwitchResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditScopeRuleSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditScopeRuleSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetAuditScopeRuleSwitchResponse", string(data)}, " ")
}
