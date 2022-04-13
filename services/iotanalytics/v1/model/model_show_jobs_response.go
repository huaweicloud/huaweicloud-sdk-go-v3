package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobsResponse struct {
	// 总数

	Count *int64 `json:"count,omitempty"`
	// 作业列表

	Jobs           *[]StreamingJobInfoDto `json:"jobs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobsResponse struct{}"
	}

	return strings.Join([]string{"ShowJobsResponse", string(data)}, " ")
}
