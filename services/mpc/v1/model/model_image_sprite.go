package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageSprite struct {
	Params *ImageSpritePara `json:"params"`

	Output *ObsObjInfo `json:"output,omitempty"`

	// 截取雪碧图后，雪碧图图片文件的输出文件名，如果不填，则默认为：{inputName}_imageSprite_{雪碧图id}_{number}.{format}.{雪碧图id}和{number}从0开始递增
	OutputObjectName *string `json:"output_object_name,omitempty"`

	// 截取雪碧图后，Web VTT 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：{inputName}_imageSprite_{雪碧图_id}.vtt
	WebvttObjectName *string `json:"webvtt_object_name,omitempty"`
}

func (o ImageSprite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSprite struct{}"
	}

	return strings.Join([]string{"ImageSprite", string(data)}, " ")
}
