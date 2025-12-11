package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditRuleSqlResponse Response Object
type DeleteAuditRuleSqlResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditRuleSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditRuleSqlResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditRuleSqlResponse", string(data)}, " ")
}
