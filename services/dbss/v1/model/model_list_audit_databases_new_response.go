package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditDatabasesNewResponse Response Object
type ListAuditDatabasesNewResponse struct {

	// 数据库信息列表
	Databases *[]DataBaseBean `json:"databases,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditDatabasesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditDatabasesNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditDatabasesNewResponse", string(data)}, " ")
}
