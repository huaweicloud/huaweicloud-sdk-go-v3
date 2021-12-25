package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTopicsResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回数量

	Size *int32 `json:"size,omitempty"`
	// 主题列表

	Items          *[]Topic `json:"items,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListTopicsResponse", string(data)}, " ")
}
