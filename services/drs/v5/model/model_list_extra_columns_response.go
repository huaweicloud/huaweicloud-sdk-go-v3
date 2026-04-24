package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExtraColumnsResponse Response Object
type ListExtraColumnsResponse struct {

	// 列表中的项目总数，与分页无关。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 列加工对象
	ColumnProcessObjects *[]ColumnProcessObjects `json:"column_process_objects,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListExtraColumnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtraColumnsResponse struct{}"
	}

	return strings.Join([]string{"ListExtraColumnsResponse", string(data)}, " ")
}
