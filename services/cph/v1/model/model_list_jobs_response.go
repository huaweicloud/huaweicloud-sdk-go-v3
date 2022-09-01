package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListJobsResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 任务信息
	Jobs           []interface{} `json:"jobs" xml:"jobs"`
	HttpStatusCode int           `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
