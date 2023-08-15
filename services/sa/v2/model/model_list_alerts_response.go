package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertsResponse Response Object
type ListAlertsResponse struct {

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

	// 告警列表
	Data *[]ListAlertDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAlertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertsResponse struct{}"
	}

	return strings.Join([]string{"ListAlertsResponse", string(data)}, " ")
}
