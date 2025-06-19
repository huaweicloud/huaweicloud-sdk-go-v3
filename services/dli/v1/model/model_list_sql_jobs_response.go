package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobsResponse Response Object
type ListSqlJobsResponse struct {

	// 参数解释:  请求发送是否成功。“true”表示请求发送成功 示例: true 约束限制:  无 取值范围: true, false 默认取值: 无
	IsSuccess *bool `json:"is_success,omitempty"`

	// 参数解释:  系统提示信息，执行成功时，信息可能为空 示例: import data to table t2 started 约束限制:  无 取值范围: 无 默认取值: 无
	Message *string `json:"message,omitempty"`

	// 参数解释:  作业总个数 示例: 7 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
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
