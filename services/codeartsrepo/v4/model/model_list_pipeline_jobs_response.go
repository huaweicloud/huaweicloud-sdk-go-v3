package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineJobsResponse Response Object
type ListPipelineJobsResponse struct {
	Body *[]JobWithStageDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPipelineJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineJobsResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineJobsResponse", string(data)}, " ")
}
