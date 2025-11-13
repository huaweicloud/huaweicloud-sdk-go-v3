package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobsResponse Response Object
type ListDbAgentJobsResponse struct {

	// 作业列表。
	Jobs *[]ListDbAgentJobsResult `json:"jobs,omitempty"`

	// 作业总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbAgentJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobsResponse struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobsResponse", string(data)}, " ")
}
