package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEditTaskRequestBody struct {

	// 编辑任务输入
	Inputs []EditMediaTaskInput `json:"inputs"`

	// 输出文件名
	FileName *string `json:"file_name,omitempty"`

	EditingSettings *EditingSetting `json:"editing_settings,omitempty"`

	// 回调地址
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 自定义上下文
	SessionContext *string `json:"session_context,omitempty"`

	Output *ObsInfo `json:"output"`

	MediaProcessTask *AdditionalObjectProcessReq `json:"media_process_task,omitempty"`
}

func (o CreateEditTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEditTaskRequestBody", string(data)}, " ")
}
