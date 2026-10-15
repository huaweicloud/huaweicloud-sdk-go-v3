package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkJobsResponse Response Object
type ListSparkJobsResponse struct {
	PageInfo *SparkMarkerPageInfo `json:"page_info,omitempty"`

	// **参数解释**：Spark作业列表，包含作业详细信息。
	Jobs           *[]ListSparkJobResponseDto `json:"jobs,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSparkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobsResponse struct{}"
	}

	return strings.Join([]string{"ListSparkJobsResponse", string(data)}, " ")
}
