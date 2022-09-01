package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群日志记录实体对象。
type ClusterLogRecord struct {

	// 日志任务ID，通过系统UUID生成。
	Id *string `json:"id,omitempty" xml:"id"`

	// 集群ID。
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId"`

	// 创建时间。格式：Unix时间戳格式。
	CreateAt *string `json:"createAt,omitempty" xml:"createAt"`

	// 日志在OBS桶中的备份路径。
	LogPath *string `json:"logPath,omitempty" xml:"logPath"`

	// 任务状态。 - RUNNING: 备份行中。 - SUCCESS: 备份成功。 - FAIL: 备份失败。
	Status *string `json:"status,omitempty" xml:"status"`

	// 结束时间，当创建未结束时结束时间为null。格式：Unix时间戳格式。
	FinishedAt *int64 `json:"finishedAt,omitempty" xml:"finishedAt"`

	// 任务类型。 - Manual: 手动备份。 - Auto： 自动备份。
	JobTypes *string `json:"jobTypes,omitempty" xml:"jobTypes"`

	// 错误信息。当任务状态没有处于失败状态时该字段为null。
	FailedMsg *string `json:"failedMsg,omitempty" xml:"failedMsg"`

	// 任务ID。
	JobId *string `json:"jobId,omitempty" xml:"jobId"`
}

func (o ClusterLogRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterLogRecord struct{}"
	}

	return strings.Join([]string{"ClusterLogRecord", string(data)}, " ")
}
