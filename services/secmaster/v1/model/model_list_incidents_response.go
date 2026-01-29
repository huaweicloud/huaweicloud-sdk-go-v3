package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentsResponse Response Object
type ListIncidentsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 事件总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Size *int32 `json:"size,omitempty"`

	// 分页的页码
	Page *int32 `json:"page,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 事件列表
	Data *[]IncidentDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIncidentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentsResponse struct{}"
	}

	return strings.Join([]string{"ListIncidentsResponse", string(data)}, " ")
}
