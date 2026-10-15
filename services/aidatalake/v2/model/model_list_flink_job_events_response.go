package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkJobEventsResponse Response Object
type ListFlinkJobEventsResponse struct {

	// **参数解释**：作业事件总个数。 **取值范围**：0~2147483647。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：Flink作业事件列表。可显示0~1000个FlinkJobEvent对象。
	JobEvents      *[]FlinkJobEvent `json:"job_events,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListFlinkJobEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobEventsResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkJobEventsResponse", string(data)}, " ")
}
