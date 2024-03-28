package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesRequest Request Object
type ListTablesRequest struct {

	// 查看表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	CurrentPage *int32 `json:"current-page,omitempty"`

	// 过滤表名称的关键词。
	Keyword *string `json:"keyword,omitempty"`

	PageSize *int32 `json:"page-size,omitempty"`

	TableType *string `json:"table-type,omitempty"`

	// 是否获取表的详细信息（所有者，size等）。
	WithDetail *bool `json:"with-detail,omitempty"`

	WithPriv *bool `json:"with-priv,omitempty"`
}

func (o ListTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesRequest struct{}"
	}

	return strings.Join([]string{"ListTablesRequest", string(data)}, " ")
}
