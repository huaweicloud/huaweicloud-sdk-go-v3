package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type CreateImageWatermarkRequestBody struct {
	// 要添加水印的图片文件，添加的图片短边尺寸需要超过512像素。
	File *def.FilePart `json:"file"`

	// 嵌入暗水印的内容，长度不超过32个字符。当前仅支持数字及英文大小写。
	BlindWatermark *def.MultiPart `json:"blind_watermark"`
}

func (o CreateImageWatermarkRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkRequestBody", string(data)}, " ")
}
