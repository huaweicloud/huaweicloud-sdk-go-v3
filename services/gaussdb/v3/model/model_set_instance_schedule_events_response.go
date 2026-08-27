package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceScheduleEventsResponse Response Object
type SetInstanceScheduleEventsResponse struct {

	// **参数解释**：  设置事件执行策略响应结果。
	Results        *[]EventJobResult `json:"results,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o SetInstanceScheduleEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceScheduleEventsResponse struct{}"
	}

	return strings.Join([]string{"SetInstanceScheduleEventsResponse", string(data)}, " ")
}
