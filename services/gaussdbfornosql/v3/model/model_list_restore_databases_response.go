package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreDatabasesResponse Response Object
type ListRestoreDatabasesResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据库名称列表。
	DatabaseNames  *[]string `json:"database_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRestoreDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreDatabasesResponse", string(data)}, " ")
}
