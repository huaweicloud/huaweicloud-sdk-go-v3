package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PutEventsResponse struct {

	// 发布失败的事件个数
	FailedCount *int32 `json:"failed_count,omitempty" xml:"failed_count"`

	// 事件信息
	Events         *[]PutEventsRespEvents `json:"events,omitempty" xml:"events"`
	HttpStatusCode int                    `json:"-"`
}

func (o PutEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsResponse struct{}"
	}

	return strings.Join([]string{"PutEventsResponse", string(data)}, " ")
}
