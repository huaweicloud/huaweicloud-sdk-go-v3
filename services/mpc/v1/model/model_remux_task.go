package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemuxTask struct {

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

	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	OutputParam *RemuxOutputParam `json:"output_param,omitempty" xml:"output_param"`

	// 任务完成进度百分比值。
	CompleteRatio *int32 `json:"complete_ratio,omitempty" xml:"complete_ratio"`

	OutputMetadata *MetaData `json:"output_metadata,omitempty" xml:"output_metadata"`
}

func (o RemuxTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemuxTask struct{}"
	}

	return strings.Join([]string{"RemuxTask", string(data)}, " ")
}
