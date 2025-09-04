package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountImageSubJobsRequest Request Object
type CountImageSubJobsRequest struct {

	// job详情的状态： * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 成功 * `FAILED` - 失败 * `ABNORMAL` - 异常 * `ROLLBACK` - 回滚中 * `ABORTING` - 取消
	Status *string `json:"status,omitempty"`

	// 任务类型，传入多个任务类型的时候将任务类型用英文逗号(,)进行分隔。
	JobType *string `json:"job_type,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o CountImageSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountImageSubJobsRequest struct{}"
	}

	return strings.Join([]string{"CountImageSubJobsRequest", string(data)}, " ")
}
