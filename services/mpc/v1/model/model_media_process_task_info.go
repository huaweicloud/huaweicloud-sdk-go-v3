package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 媒体处理信息
type MediaProcessTaskInfo struct {
	// 任务Id

	TaskId *string `json:"task_id,omitempty"`
	// 任务执行状态，取值如下。 \"NO_TASK\"      //无任务，task_id非法 \"WAITING\"      //等待启动 \"PROCESSING\"   //处理中 \"SUCCEEDED\"    //成功 \"FAILED\"       //失败 \"CANCELED\"     //已删除

	Status *string `json:"status,omitempty"`
	// 任务启动时间

	CreateTime *string `json:"create_time,omitempty"`
	// 任务结束时间

	EndTime *string `json:"end_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`
	// 任务描述，当出现异常时，此字段为异常的原因。

	Description *string `json:"description,omitempty"`
	// 输出文件名。

	OutputFileName *[]string `json:"output_file_name,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`
}

func (o MediaProcessTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaProcessTaskInfo struct{}"
	}

	return strings.Join([]string{"MediaProcessTaskInfo", string(data)}, " ")
}
