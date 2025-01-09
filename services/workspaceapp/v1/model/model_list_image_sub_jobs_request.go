package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageSubJobsRequest Request Object
type ListImageSubJobsRequest struct {

	// job详情的状态： * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 成功 * `FAILED` - 失败 * `ABNORMAL` - 异常 * `ROLLBACK` - 回滚中 * `ABORTING` - 取消
	Status *string `json:"status,omitempty"`

	// job类型： * `CREATE_SERVER` - 创建镜像实例 * `CREATE_SERVER_IMAGE` - 构建镜像 * `DELETE_SERVER` - 删除镜像实例
	JobType string `json:"job_type"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ListImageSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageSubJobsRequest struct{}"
	}

	return strings.Join([]string{"ListImageSubJobsRequest", string(data)}, " ")
}
