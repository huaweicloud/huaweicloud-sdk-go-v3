package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExecuteEventsResponse Response Object
type BatchExecuteEventsResponse struct {

	// **参数解释**: 事件操作响应结果。
	Results        *[]EventJobResult `json:"results,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchExecuteEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteEventsResponse struct{}"
	}

	return strings.Join([]string{"BatchExecuteEventsResponse", string(data)}, " ")
}
