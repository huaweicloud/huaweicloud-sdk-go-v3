package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSqlsResponse Response Object
type ListAuditSqlsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 总数
	Count *int32 `json:"count,omitempty"`

	// sql语句列表
	Sqls           *[]AuditSqlResponseSqls `json:"sqls,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAuditSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditSqlsResponse", string(data)}, " ")
}
