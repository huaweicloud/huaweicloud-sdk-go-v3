package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPictureModelingJobsResponse Response Object
type ListPictureModelingJobsResponse struct {

	// 照片建模任务总数。
	Count *int32 `json:"count,omitempty"`

	// 照片建模任务列表。
	PictureModelingJobs *[]PictureModelingInfo `json:"picture_modeling_jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPictureModelingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPictureModelingJobsResponse struct{}"
	}

	return strings.Join([]string{"ListPictureModelingJobsResponse", string(data)}, " ")
}
