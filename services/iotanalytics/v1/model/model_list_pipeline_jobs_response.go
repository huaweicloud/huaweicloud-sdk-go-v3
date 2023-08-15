package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineJobsResponse Response Object
type ListPipelineJobsResponse struct {

	// 总数
	Count *int64 `json:"count,omitempty"`

	// 管道列表
	Pipelines      *[]PipelineJobInfoDto `json:"pipelines,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListPipelineJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineJobsResponse struct{}"
	}

	return strings.Join([]string{"ListPipelineJobsResponse", string(data)}, " ")
}
