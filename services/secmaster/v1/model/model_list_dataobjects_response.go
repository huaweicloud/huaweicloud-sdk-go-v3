package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataobjectsResponse Response Object
type ListDataobjectsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Size *int32 `json:"size,omitempty"`

	// 分页的页码
	Page *int32 `json:"page,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 详情列表
	Data           *[]DataObjectDetail `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListDataobjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataobjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDataobjectsResponse", string(data)}, " ")
}
