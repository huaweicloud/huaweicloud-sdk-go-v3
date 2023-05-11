package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryEditingJob struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 任务执行状态，取值如下。WAITING：等待启动 PROCESSING：处理中 SUCCEEDED：处理成功。FAILED：处理失败。
	Status *string `json:"status,omitempty"`

	// 任务的返回码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 处理信息。
	ExecDescription *string `json:"exec_description,omitempty"`

	// 处理信息。
	Description *string `json:"description,omitempty"`

	// 任务执行进度百分比, 取值范围：[0, 100]。
	Progress *int32 `json:"progress,omitempty"`

	// 剪切拼接任务创建时间 格式 yyyyMMddHHmmss 必须是与时区无关的UTC时间
	CreateTime *string `json:"create_time,omitempty"`

	// 剪切拼接任务开始时间 格式 yyyyMMddHHmmss 必须是与时区无关的UTC时间
	StartTime *string `json:"start_time,omitempty"`

	// 剪切拼接任务结束时间 格式 yyyyMMddHHmmss 必须是与时区无关的UTC时间
	EndTime *string `json:"end_time,omitempty"`

	// 用户自定义数据
	UserData *string `json:"user_data,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	EditingOutput *EditingOutputFileInfo `json:"editing_output,omitempty"`
}

func (o QueryEditingJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryEditingJob struct{}"
	}

	return strings.Join([]string{"QueryEditingJob", string(data)}, " ")
}
