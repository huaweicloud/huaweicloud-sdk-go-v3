package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListItaSubJobsRequest struct {

	// 任务状态 - SUCCESS：成功。 - RUNNING：运行中。 - FAILED：失败。 - WAITING：等待。
	Status *string `json:"status,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务类型  - createDesktops：创建桌面任务。  - applyWorkspace：开通云桌面服务。  - cancelWorkspace：注销云桌面服务。  - expandVolumes:  扩容磁盘。  - addVolumes: 添加磁盘。
	JobType *string `json:"job_type,omitempty"`

	// 用于分页查询，取值范围0~1000，默认1000。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListItaSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListItaSubJobsRequest struct{}"
	}

	return strings.Join([]string{"ListItaSubJobsRequest", string(data)}, " ")
}
