package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobsResponse Response Object
type ListSqlJobsResponse struct {

	// 请求发送是否成功。“true”表示请求发送成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 作业总个数。
	JobCount *int32 `json:"job_count,omitempty"`

	// 作业信息。
	Jobs           *[]SqlJob `json:"jobs,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSqlJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobsResponse struct{}"
	}

	return strings.Join([]string{"ListSqlJobsResponse", string(data)}, " ")
}
