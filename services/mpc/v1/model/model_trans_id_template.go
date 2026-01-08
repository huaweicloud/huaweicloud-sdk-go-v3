package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransIdTemplate struct {

	// 输出视频对应的模板ID
	TemplateId int32 `json:"template_id"`

	Output *ObsObjInfo `json:"output,omitempty"`

	// 输出文件名
	OutputFilename *string `json:"output_filename,omitempty"`
}

func (o TransIdTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransIdTemplate struct{}"
	}

	return strings.Join([]string{"TransIdTemplate", string(data)}, " ")
}
