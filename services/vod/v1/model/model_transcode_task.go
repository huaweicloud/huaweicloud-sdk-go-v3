package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeTask struct {

	// 转码模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 任务状态，包含成功，失败
	Status *string `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 转码输出对象信息
	Output *[]TaskOutPut `json:"output,omitempty"`
}

func (o TranscodeTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeTask struct{}"
	}

	return strings.Join([]string{"TranscodeTask", string(data)}, " ")
}
