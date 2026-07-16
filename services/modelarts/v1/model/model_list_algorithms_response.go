package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlgorithmsResponse Response Object
type ListAlgorithmsResponse struct {

	// 查询到当前用户名下的所有算法总数。
	Total *int32 `json:"total,omitempty"`

	// 查询到当前用户名下的所有符合查询条件的算法总数。
	Count *int32 `json:"count,omitempty"`

	// 查询到当前用户名下的所有算法限制个数。
	Limit *int32 `json:"limit,omitempty"`

	// 查询到当前用户名下的所有算法查询偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询到当前用户名下的所有算法排序依赖字段。
	SortBy *string `json:"sort_by,omitempty"`

	// 查询到当前用户名下的所有算法排序方式，默认为“desc”，降序排序。也可以选择对应的“asc”，升序排序。
	Order *string `json:"order,omitempty"`

	// 查询到当前用户名下的所有算法分组方式。
	GroupBy *string `json:"group_by,omitempty"`

	// 查询到当前用户名下的所有符合查询条件的算法详情。
	Items          *[]AlgorithmResponse `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAlgorithmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlgorithmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlgorithmsResponse", string(data)}, " ")
}
