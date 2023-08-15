package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataTablesResponse Response Object
type ListDataTablesResponse struct {

	// 当前数据库中表的记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 表的列表
	Tables         *[]TablesList `json:"tables,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDataTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataTablesResponse struct{}"
	}

	return strings.Join([]string{"ListDataTablesResponse", string(data)}, " ")
}
