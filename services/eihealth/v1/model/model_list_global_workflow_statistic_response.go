package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalWorkflowStatisticResponse Response Object
type ListGlobalWorkflowStatisticResponse struct {

	// 所有作业总数
	JobCount       *int32 `json:"job_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGlobalWorkflowStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalWorkflowStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalWorkflowStatisticResponse", string(data)}, " ")
}
