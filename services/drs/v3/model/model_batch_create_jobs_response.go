package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateJobsResponse struct {
	// 批量创建任务的响应体集合。

	Results *[]CreateJobResp `json:"results,omitempty"`
	// 总记录数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateJobsResponse", string(data)}, " ")
}
