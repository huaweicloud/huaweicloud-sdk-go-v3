package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTablesRequest struct {
	// 过滤表名称的关键词。

	Keyword *string `json:"keyword,omitempty"`
	// 过滤标签的关键字

	Tag *string `json:"tag,omitempty"`
	// 当前偏移量，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的最大作业个数，范围: [1, 100]。默认值：10。

	Limit *int32 `json:"limit,omitempty"`
	// 指定作业排序字段，默认为created_time（作业创建时间），支持created_time(作业创建时间)、modified_time（作业更新时间） 、job_name（作业名称）三种排序字段。

	OrderBy *string `json:"order_by,omitempty"`
	// 指定作业排序的升降序，默认为desc（降序），支持asc（升序）、desc（降序）两种排序方式。

	Order *string `json:"order,omitempty"`
}

func (o ListTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesRequest struct{}"
	}

	return strings.Join([]string{"ListTablesRequest", string(data)}, " ")
}
