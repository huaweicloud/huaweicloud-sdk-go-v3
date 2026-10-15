package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkJobAutoscalingHistoriesResponse Response Object
type ListFlinkJobAutoscalingHistoriesResponse struct {

	// **参数解释**：调优历史记录列表。
	AutoscalingHistories *[]AutoscalingHistory `json:"autoscaling_histories,omitempty"`

	// **参数解释**：调优历史记录总数。 **取值范围**：不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFlinkJobAutoscalingHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobAutoscalingHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkJobAutoscalingHistoriesResponse", string(data)}, " ")
}
