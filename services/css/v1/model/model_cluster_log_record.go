package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群日志记录实体对象。
type ClusterLogRecord struct {
	// 日志任务ID，通过系统uuid生成。

	Id *string `json:"id,omitempty"`
	// 集群ID。

	ClusterId *string `json:"clusterId,omitempty"`
	// 创建时间。

	CreateAt *string `json:"createAt,omitempty"`
	// 备份路径。

	LogPath *string `json:"logPath,omitempty"`
	// 任务状态。

	Status *string `json:"status,omitempty"`
	// 结束时间。

	FinishedAt *int64 `json:"finishedAt,omitempty"`
	// 任务类型。

	JobTypes *string `json:"jobTypes,omitempty"`
	// 错误信息。

	FailedMsg *string `json:"failedMsg,omitempty"`
	// 任务ID。

	JobId *string `json:"jobId,omitempty"`
}

func (o ClusterLogRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterLogRecord struct{}"
	}

	return strings.Join([]string{"ClusterLogRecord", string(data)}, " ")
}
