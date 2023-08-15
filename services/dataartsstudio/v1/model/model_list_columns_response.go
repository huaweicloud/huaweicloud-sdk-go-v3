package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListColumnsResponse Response Object
type ListColumnsResponse struct {

	// 表id
	TableId *string `json:"table_id,omitempty"`

	// 当前表中字段记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 字段列表
	Columns        *[]ColumnsList `json:"columns,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListColumnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListColumnsResponse struct{}"
	}

	return strings.Join([]string{"ListColumnsResponse", string(data)}, " ")
}
