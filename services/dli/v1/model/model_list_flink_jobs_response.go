package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkJobsResponse Response Object
type ListFlinkJobsResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *string `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	JobList        *FlinkJobList `json:"job_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsResponse", string(data)}, " ")
}
