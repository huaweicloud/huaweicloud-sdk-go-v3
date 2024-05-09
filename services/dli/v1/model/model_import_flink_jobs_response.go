package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobsResponse Response Object
type ImportFlinkJobsResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *string `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	// 作业导入结果。
	JobMapping     *[]JobMap `json:"job_mapping,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ImportFlinkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobsResponse struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobsResponse", string(data)}, " ")
}
