package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneHpcCacheTaskInfoResp 每个任务信息
type OneHpcCacheTaskInfoResp struct {

	// 任务id
	TaskId string `json:"task_id"`

	// 任务类型
	Type string `json:"type"`

	// 任务状态
	Status string `json:"status"`

	// 源端对象
	SrcTarget string `json:"src_target"`

	// 源端路径
	SrcPrefix string `json:"src_prefix"`

	// 目的端对象
	DestTarget string `json:"dest_target"`

	// 目的端路径
	DestPrefix string `json:"dest_prefix"`

	// 任务开始时间
	StartTime string `json:"start_time"`

	// 任务结束时间
	EndTime string `json:"end_time"`

	// 任务执行结果信息
	Message string `json:"message"`
}

func (o OneHpcCacheTaskInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneHpcCacheTaskInfoResp struct{}"
	}

	return strings.Join([]string{"OneHpcCacheTaskInfoResp", string(data)}, " ")
}
