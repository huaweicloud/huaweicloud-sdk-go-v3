package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentVideoImageDetail struct {

	// 截帧在视频文件中的时间，单位为秒
	Time *float32 `json:"time,omitempty"`

	// 截帧审核结果是否通过。 block：包含敏感信息，不通过 review：需要人工复检
	Suggestion *string `json:"suggestion,omitempty"`

	// 截帧检测出的文本
	OcrText *string `json:"ocr_text,omitempty"`

	// 识别的详细标签
	Label *string `json:"label,omitempty"`

	// 命中的文本风险片段列表
	Segments *[]DocumentVideoImageDetailSegments `json:"segments,omitempty"`
}

func (o DocumentVideoImageDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentVideoImageDetail struct{}"
	}

	return strings.Join([]string{"DocumentVideoImageDetail", string(data)}, " ")
}
