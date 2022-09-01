package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditingJob struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务状态。  取值如下： - INIT：初始状态。 - WAITING：等待启动。 - PROCESSING：处理中。 - SUCCEED：处理成功。 - FAILED：处理失败。 - CANCELED：已取消。
	Status *string `json:"status,omitempty" xml:"status"`

	// 任务创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 任务启动时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 任务的返回码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 用户数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 剪辑任务类型。取值如下：\"CLIP\",\"CONCAT\",\"CONCATS\",\"MIX\"。
	EditType *[]string `json:"edit_type,omitempty" xml:"edit_type"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	EditTaskReq *CreateEditingJobReq `json:"edit_task_req,omitempty" xml:"edit_task_req"`

	// 剪辑输出meta信息
	OutputFileInfo *[]OutputFileInfo `json:"output_file_info,omitempty" xml:"output_file_info"`
}

func (o EditingJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingJob struct{}"
	}

	return strings.Join([]string{"EditingJob", string(data)}, " ")
}
