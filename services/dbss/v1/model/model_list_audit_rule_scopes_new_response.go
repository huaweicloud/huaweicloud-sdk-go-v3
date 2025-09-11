package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleScopesNewResponse Response Object
type ListAuditRuleScopesNewResponse struct {

	// 审计范围规则列表
	Scopes *[]RuleScopeInfo `json:"scopes,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditRuleScopesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleScopesNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditRuleScopesNewResponse", string(data)}, " ")
}
