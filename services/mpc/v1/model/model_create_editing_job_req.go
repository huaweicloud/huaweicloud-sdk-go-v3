package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建剪辑任务
type CreateEditingJobReq struct {
	// 剪辑任务类型。取值如下：\"CLIP\",\"CONCAT\",\"CONCATS\",\"MIX\"。

	EditType *[]string `json:"edit_type,omitempty"`
	// 剪切信息

	Clips *[]ClipInfo `json:"clips,omitempty"`
	// 多拼接任务信息，支持多个拼接输出，与concat参数只能二选一。

	Concats *[]MultiConcatInfo `json:"concats,omitempty"`

	Concat *ConcatInfo `json:"concat,omitempty"`

	Mix *MixInfo `json:"mix,omitempty"`

	OutputSetting *OutputSetting `json:"output_setting,omitempty"`
	// 水印信息。

	ImageWatermarkSettings *[]ImageWatermarkSetting `json:"image_watermark_settings,omitempty"`
	// 用户自定义数据。

	UserData *string `json:"user_data,omitempty"`
}

func (o CreateEditingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobReq struct{}"
	}

	return strings.Join([]string{"CreateEditingJobReq", string(data)}, " ")
}
