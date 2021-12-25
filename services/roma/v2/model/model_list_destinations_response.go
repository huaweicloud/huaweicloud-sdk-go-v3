package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDestinationsResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回数量

	Size *int32 `json:"size,omitempty"`
	// 目标数据源列表

	Items          *[]Destination `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListDestinationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDestinationsResponse struct{}"
	}

	return strings.Join([]string{"ListDestinationsResponse", string(data)}, " ")
}
