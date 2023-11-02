package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateItemResp 集群升级路径响应体
type UpdateItemResp struct {

	// 升级项ID
	Id *string `json:"id,omitempty"`

	// 起始版本
	From *string `json:"from,omitempty"`

	// 目标版本
	To *string `json:"to,omitempty"`

	// 升级路径状态
	Status *string `json:"status,omitempty"`

	// 升级进度
	Process *string `json:"process,omitempty"`

	// 起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 升级任务ID
	JobId *string `json:"job_id,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o UpdateItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateItemResp struct{}"
	}

	return strings.Join([]string{"UpdateItemResp", string(data)}, " ")
}
