package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkJobsResponse Response Object
type ListFlinkJobsResponse struct {

	// 参数解释:  执行请求是否成功。“true”表示请求执行成功 示例: true 约束限制:  无 取值范围: true,false 默认取值: 无
	IsSuccess *string `json:"is_success,omitempty"`

	// 参数解释:  系统提示信息，执行成功时，信息可能为空 示例: success 约束限制:  无 取值范围: 无 默认取值: 无
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
