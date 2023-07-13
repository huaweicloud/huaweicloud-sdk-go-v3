package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutOfficialEventsResponse Response Object
type PutOfficialEventsResponse struct {

	// 发布失败的事件个数
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 事件列表
	Events *[]PutEventsRespEvents `json:"events,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PutOfficialEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutOfficialEventsResponse struct{}"
	}

	return strings.Join([]string{"PutOfficialEventsResponse", string(data)}, " ")
}
