package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartJobSubmission struct {

	// 作业是否为增量迁移
	IsIncrementing *bool `json:"isIncrementing,omitempty" xml:"isIncrementing"`

	// 删除数据行数
	DeleteRows *int32 `json:"delete_rows,omitempty" xml:"delete_rows"`

	// 更新数据行数
	UpdateRows *int32 `json:"update_rows,omitempty" xml:"update_rows"`

	// 写入数据行数
	WriteRows *int32 `json:"write_rows,omitempty" xml:"write_rows"`

	// 作业提交id
	SubmissionId *int32 `json:"submission-id,omitempty" xml:"submission-id"`

	// 作业名称
	JobName string `json:"job-name" xml:"job-name"`

	// 创建用户
	CreationUser string `json:"creation-user" xml:"creation-user"`

	// 创建时间，单位：毫秒。
	CreationDate int64 `json:"creation-date" xml:"creation-date"`

	// 执行时间
	ExecuteDate *int64 `json:"execute-date,omitempty" xml:"execute-date"`

	// 作业进度，失败时为“-1”，其它情况为0～100
	Progress float32 `json:"progress" xml:"progress"`

	// 作业状态： - BOOTING：启动中。 - FAILURE_ON_SUBMIT：提交失败。 - RUNNING：运行中。 - SUCCEEDED：成功。 - FAILED：失败。 - UNKNOWN：未知。 - NEVER_EXECUTED：未被执行
	Status string `json:"status" xml:"status"`

	// 是否停止增量迁移
	IsStopingIncrement *string `json:"isStopingIncrement,omitempty" xml:"isStopingIncrement"`

	// 是否定时执行作业
	IsExecuteAuto *bool `json:"is-execute-auto,omitempty" xml:"is-execute-auto"`

	// 作业最后更新时间
	LastUpdateDate *int64 `json:"last-update-date,omitempty" xml:"last-update-date"`

	// 最后更新作业状态的用户
	LastUdpateUser *string `json:"last-udpate-user,omitempty" xml:"last-udpate-user"`

	// 作业执行完成后是否删除
	IsDeleteJob *bool `json:"isDeleteJob,omitempty" xml:"isDeleteJob"`
}

func (o StartJobSubmission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobSubmission struct{}"
	}

	return strings.Join([]string{"StartJobSubmission", string(data)}, " ")
}
