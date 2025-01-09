package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageJobsRequest Request Object
type ListImageJobsRequest struct {

	// job状态： * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 完成 * `FAILED` - 失败
	Status *string `json:"status,omitempty"`

	// job类型： * `CREATE_SERVER` - 创建镜像实例 * `CREATE_SERVER_IMAGE` - 构建镜像 * `DELETE_SERVER` - 删除镜像实例
	JobType string `json:"job_type"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListImageJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageJobsRequest struct{}"
	}

	return strings.Join([]string{"ListImageJobsRequest", string(data)}, " ")
}
