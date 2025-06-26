package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubtaskDetail struct {

	// 子任务ID
	Id int32 `json:"id"`

	// 内部任务ID
	JobId int32 `json:"job_id"`

	// 命名空间名
	Namespace string `json:"namespace"`

	// 制品仓库
	Repository string `json:"repository"`

	// 制品版本
	Tag string `json:"tag"`

	// 摘要
	Digest string `json:"digest"`

	// 状态
	Status string `json:"status"`

	// 状态信息
	StatusText string `json:"status_text"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`
}

func (o SubtaskDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtaskDetail struct{}"
	}

	return strings.Join([]string{"SubtaskDetail", string(data)}, " ")
}
