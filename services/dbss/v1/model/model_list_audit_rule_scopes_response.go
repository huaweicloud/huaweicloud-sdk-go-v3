package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleScopesResponse Response Object
type ListAuditRuleScopesResponse struct {

	// 审计范围规则列表
	Scopes *[]RuleScopeInfo `json:"scopes,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditRuleScopesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleScopesResponse struct{}"
	}

	return strings.Join([]string{"ListAuditRuleScopesResponse", string(data)}, " ")
}
