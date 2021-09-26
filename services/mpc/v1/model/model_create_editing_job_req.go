package model

import (
	"encoding/json"

	"strings"
)

// 创建剪辑任务
type CreateEditingJobReq struct {
	// 剪辑任务类型。取值如下：\"CLIP\",\"CONCAT\"。

	EditType *[]string `json:"edit_type,omitempty"`
	// 剪切信息

	Clips *[]ClipInfo `json:"clips,omitempty"`

	Contcat *ConcatInfo `json:"contcat,omitempty"`

	OutputSetting *OutputSetting `json:"output_setting,omitempty"`
	// 水印信息。

	ImageWatermarkSettings *[]ImageWatermarkSetting `json:"image_watermark_settings,omitempty"`
	// 用户自定义数据。

	UserData *string `json:"user_data,omitempty"`
}

func (o CreateEditingJobReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEditingJobReq struct{}"
	}

	return strings.Join([]string{"CreateEditingJobReq", string(data)}, " ")
}
