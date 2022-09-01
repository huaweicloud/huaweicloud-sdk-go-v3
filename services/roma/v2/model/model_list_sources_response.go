package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSourcesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 数据源列表
	Items          *[]Source `json:"items,omitempty" xml:"items"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourcesResponse struct{}"
	}

	return strings.Join([]string{"ListSourcesResponse", string(data)}, " ")
}
