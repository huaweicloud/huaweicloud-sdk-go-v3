package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQueuesResponse struct {
	// 该租户的所有队列总数。

	Total *int32 `json:"total,omitempty"`
	// 该租户的所有队列数组。

	Queues         *[]ListQueuesRespQueues `json:"queues,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListQueuesResponse", string(data)}, " ")
}
