package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartJobSubmission struct {
	// 作业是否为增量迁移

	IsIncrementing *bool `json:"isIncrementing,omitempty"`
	// 删除数据行数

	DeleteRows *int32 `json:"delete_rows,omitempty"`
	// 更新数据行数

	UpdateRows *int32 `json:"update_rows,omitempty"`
	// 写入数据行数

	WriteRows *int32 `json:"write_rows,omitempty"`
	// 作业提交id

	SubmissionId *int32 `json:"submission-id,omitempty"`
	// 作业名称

	JobName string `json:"job-name"`
	// 创建用户

	CreationUser string `json:"creation-user"`
	// 创建时间，单位：毫秒。

	CreationDate int64 `json:"creation-date"`
	// 执行时间

	ExecuteDate *int64 `json:"execute-date,omitempty"`
	// 作业进度，失败时为“-1”，其它情况为0～100

	Progress int32 `json:"progress"`
	// 作业状态： - BOOTING：启动中。 - FAILURE_ON_SUBMIT：提交失败。 - RUNNING：运行中。 - SUCCEEDED：成功。 - FAILED：失败。 - UNKNOWN：未知。 - NEVER_EXECUTED：未被执行

	Status string `json:"status"`
	// 是否停止增量迁移

	IsStopingIncrement *bool `json:"isStopingIncrement,omitempty"`
	// 是否定时执行作业

	IsExecuteAuto *bool `json:"is-execute-auto,omitempty"`
	// 作业最后更新时间

	LastUpdateDate *int64 `json:"last-update-date,omitempty"`
	// 最后更新作业状态的用户

	LastUdpateUser *string `json:"last-udpate-user,omitempty"`
	// 作业执行完成后是否删除

	IsDeleteJob *bool `json:"isDeleteJob,omitempty"`
}

func (o StartJobSubmission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobSubmission struct{}"
	}

	return strings.Join([]string{"StartJobSubmission", string(data)}, " ")
}
