package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditDatabasesResponse Response Object
type ListAuditDatabasesResponse struct {

	// 数据库信息列表
	Databases *[]DataBaseBean `json:"databases,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListAuditDatabasesResponse", string(data)}, " ")
}
