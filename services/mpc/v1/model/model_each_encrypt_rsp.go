package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EachEncryptRsp struct {

	// 任务Id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务执行状态。  取值如下： - NO_TASK：无任务 - WAITING：等待启动 - PROCESSING：加密中 - SUCCEEDED：加密成功 - FAILED：加密失败 - CANCELED：已删除
	Status *string `json:"status,omitempty" xml:"status"`

	// 加密任务启动时间。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 加密任务结束时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 加密生成的文件名，数组类型，可能包含多个，包含加密文件名。
	OutputFileName *[]string `json:"output_file_name,omitempty" xml:"output_file_name"`

	// 用户数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 加密任务描述，当加密出现异常时，此字段为异常的原因。
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o EachEncryptRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EachEncryptRsp struct{}"
	}

	return strings.Join([]string{"EachEncryptRsp", string(data)}, " ")
}
