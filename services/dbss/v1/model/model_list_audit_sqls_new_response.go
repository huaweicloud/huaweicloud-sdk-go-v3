package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSqlsNewResponse Response Object
type ListAuditSqlsNewResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 总数
	Count *int32 `json:"count,omitempty"`

	// sql语句列表
	Sqls           *[]AuditSqlResponseSqls `json:"sqls,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAuditSqlsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSqlsNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditSqlsNewResponse", string(data)}, " ")
}
