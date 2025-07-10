package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 安装实例的用户。
	Target *string `json:"target,omitempty"`

	// 任务状态： * `INIT` - 初始化中 * `WAITING` - 等待安装结束 * `SUCCESS` - 成功 * `FAIL` - 失败任务状态
	JobStatus *string `json:"job_status,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
