package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Submission struct {

	// 作业是否为增量迁移
	IsIncrementing bool `json:"isIncrementing"`

	// 作业名称
	JobName string `json:"job-name"`

	Counters *Counters `json:"counters,omitempty"`

	// 是否停止增量迁移
	IsStopingIncrement string `json:"isStopingIncrement"`

	// 是否定时执行作业
	IsExecuteAuto bool `json:"is-execute-auto"`

	// 作业最后更新时间
	LastUpdateDate int64 `json:"last-update-date"`

	// 最后更新作业状态的用户
	LastUdpateUser string `json:"last-udpate-user"`

	// 作业执行完成后是否删除
	IsDeleteJob bool `json:"isDeleteJob"`

	// 创建用户
	CreationUser string `json:"creation-user"`

	// 创建时间
	CreationDate int64 `json:"creation-date"`

	// 作业ID
	ExternalId string `json:"external-id"`

	// 作业进度，失败时为“-1”，其它情况为0～100
	Progress float32 `json:"progress"`

	// 作业提交id
	SubmissionId *int32 `json:"submission-id,omitempty"`

	// 删除数据行数
	DeleteRows *int32 `json:"delete_rows,omitempty"`

	// 更新数据行数
	UpdateRows *int32 `json:"update_rows,omitempty"`

	// 写入数据行数
	WriteRows *int32 `json:"write_rows,omitempty"`

	// 执行时间
	ExecuteDate *int64 `json:"execute-date,omitempty"`

	// 作业状态： - BOOTING：启动中。 - FAILURE_ON_SUBMIT：提交失败。 - RUNNING：运行中。 - SUCCEEDED：成功。 - FAILED：失败。 - UNKNOWN：未知。 - NEVER_EXECUTED：未被执行
	Status string `json:"status"`

	// 错误详情，当“status”为“FAILED”时才有此字段。
	ErrorDetails *string `json:"error-details,omitempty"`

	// 错误总结，当“status”为“FAILED”时才有此字段。
	ErrorSummary *string `json:"error-summary,omitempty"`
}

func (o Submission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Submission struct{}"
	}

	return strings.Join([]string{"Submission", string(data)}, " ")
}
