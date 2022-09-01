package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 此参数在请求实体中，采用json字符串格式
type QueryLtsStructLogParams struct {

	// 搜索起始时间（UTC时间，毫秒级）。
	StartTime string `json:"start_time" xml:"start_time"`

	// 搜索结束时间（UTC时间，毫秒级）。
	EndTime string `json:"end_time" xml:"end_time"`

	// 支持SQL语句搜索， 目前支持\"GROUP BY\", \"LIKE\"和\"WHERE\"。
	SqlExpression *string `json:"sql_expression,omitempty" xml:"sql_expression"`

	// 返回内容中是否包含原始日志， 默认为false。
	OriginalContent *bool `json:"original_content,omitempty" xml:"original_content"`
}

func (o QueryLtsStructLogParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLtsStructLogParams struct{}"
	}

	return strings.Join([]string{"QueryLtsStructLogParams", string(data)}, " ")
}
