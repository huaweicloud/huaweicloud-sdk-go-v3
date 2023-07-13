package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataconnectionsResponse Response Object
type ListDataconnectionsResponse struct {

	// 当前分页返回记录数
	Count *int32 `json:"count,omitempty"`

	// 返回记录总数，一个工作空间最多只能创建50条数据连接
	MaxRecords *int32 `json:"max_records,omitempty"`

	// 返回数据连接列表
	DataConnectionLists *[]ApigDataSourceView `json:"data_connection_lists,omitempty"`
	HttpStatusCode      int                   `json:"-"`
}

func (o ListDataconnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataconnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListDataconnectionsResponse", string(data)}, " ")
}
