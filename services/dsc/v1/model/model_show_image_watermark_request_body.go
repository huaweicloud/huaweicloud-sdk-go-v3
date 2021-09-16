package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type ShowImageWatermarkRequestBody struct {
	// 待提取暗水印的图片文件。
	File *def.FilePart `json:"file"`

	// 指定待提取水印的长度，mark_len长度大于0，小于32。设置后可以提升水印提取性能
	MarkLen *def.MultiPart `json:"mark_len,omitempty"`
}

func (o ShowImageWatermarkRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkRequestBody", string(data)}, " ")
}
