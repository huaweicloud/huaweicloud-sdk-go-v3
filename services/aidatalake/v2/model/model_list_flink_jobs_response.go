package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkJobsResponse Response Object
type ListFlinkJobsResponse struct {
	PageInfo *MarkerPageInfo `json:"page_info,omitempty"`

	// **参数解释**：详细的Flink作业列表，包含0~1000个FlinkJobSummary对象。
	Jobs           *[]FlinkJobSummary `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsResponse", string(data)}, " ")
}
