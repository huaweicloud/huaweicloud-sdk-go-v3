package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunDetail struct {

	// 作业运行详情ID。
	DetailId string `json:"detail_id"`

	// 此作业的当前状态，包含提交（LAUNCHING）、运行中（RUNNING）、完成（FINISHED）、失败（FAILED）、取消（CANCELLED）。
	Status *string `json:"status,omitempty"`

	SqlJob *SqlJobRunDetail `json:"sql_job,omitempty"`
}

func (o RunDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDetail struct{}"
	}

	return strings.Join([]string{"RunDetail", string(data)}, " ")
}
