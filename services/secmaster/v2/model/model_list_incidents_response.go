package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentsResponse Response Object
type ListIncidentsResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	// tatal count
	Total *int32 `json:"total,omitempty"`

	// 当前页大小
	Limit *int32 `json:"limit,omitempty"`

	// 当前页码
	Offset *int32 `json:"offset,omitempty"`

	// success
	Success *bool `json:"success,omitempty"`

	// 事件列表
	Data *[]ListIncidentDetail `json:"data,omitempty"`

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
