package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditSqlRuleSwitchResponse Response Object
type SetAuditSqlRuleSwitchResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditSqlRuleSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditSqlRuleSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetAuditSqlRuleSwitchResponse", string(data)}, " ")
}
