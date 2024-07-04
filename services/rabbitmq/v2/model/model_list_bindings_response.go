package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBindingsResponse Response Object
type ListBindingsResponse struct {

	// 当前显示数量
	Size *int32 `json:"size,omitempty"`

	// 查询结果总数
	Total *int32 `json:"total,omitempty"`

	// 绑定信息列表
	Items          *[]BindingsDetails `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBindingsResponse struct{}"
	}

	return strings.Join([]string{"ListBindingsResponse", string(data)}, " ")
}
