package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RealTimeLogCollect struct {

	// 日志采集ID，通过系统UUID生成。
	Id *string `json:"id,omitempty"`

	// 集群ID。
	ClusterId *string `json:"clusterId,omitempty"`

	// 日志保存索引的前缀。
	IndexPrefix string `json:"indexPrefix"`

	// 日志保存时间。
	KeepDays int32 `json:"keepDays"`

	// 保存日志的目标集群ID。
	TargetClusterId string `json:"targetClusterId"`

	// 日志实时采集任务状态。
	Status string `json:"status"`

	// 日志实时采集任务开启时间。
	CreateAt int64 `json:"createAt"`

	// 日志实时采集任务更新时间。
	UpdateAt int64 `json:"updateAt"`
}

func (o RealTimeLogCollect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeLogCollect struct{}"
	}

	return strings.Join([]string{"RealTimeLogCollect", string(data)}, " ")
}
