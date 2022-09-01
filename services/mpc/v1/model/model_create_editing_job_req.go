package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建剪辑任务
type CreateEditingJobReq struct {

	// 剪辑任务类型。取值如下：\"CLIP\",\"CONCAT\",\"CONCATS\",\"MIX\"。
	EditType *[]string `json:"edit_type,omitempty" xml:"edit_type"`

	// 剪切信息
	Clips *[]ClipInfo `json:"clips,omitempty" xml:"clips"`

	// 多拼接任务信息，支持多个拼接输出，与concat参数只能二选一。
	Concats *[]MultiConcatInfo `json:"concats,omitempty" xml:"concats"`

	Concat *ConcatInfo `json:"concat,omitempty" xml:"concat"`

	Mix *MixInfo `json:"mix,omitempty" xml:"mix"`

	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	OutputSetting *OutputSetting `json:"output_setting,omitempty" xml:"output_setting"`

	// 水印信息。
	ImageWatermarkSettings *[]ImageWatermarkSetting `json:"image_watermark_settings,omitempty" xml:"image_watermark_settings"`

	// 媒体处理配置，当edit_type为空时该参数生效。会根据该参数配置，对input参数指定的源文件进行处理
	EditSettings *[]EditSetting `json:"edit_settings,omitempty" xml:"edit_settings"`

	// 用户自定义数据。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`
}

func (o CreateEditingJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEditingJobReq struct{}"
	}

	return strings.Join([]string{"CreateEditingJobReq", string(data)}, " ")
}
