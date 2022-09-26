package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageDetectionResult struct {

	// 审核结果是否通过。  block：包含敏感信息，不通过  pass：不包含敏感信息，通过  review：需要人工复检
	Suggestion *string `json:"suggestion,omitempty"`

	// 检测结果的一级标签。 支持category列表如下： politics: 涉政 terrorism: 暴恐 porn: 色情 image_text: 图文审核
	Category *string `json:"category,omitempty"`

	// 检测详情
	Details *[]ImageDetectionResultDetail `json:"details,omitempty"`

	// 图文审核检测出的文本，只有在category参数配置image_text且检测出文本时展示该字段。
	OcrText *string `json:"ocr_text,omitempty"`
}

func (o ImageDetectionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResult struct{}"
	}

	return strings.Join([]string{"ImageDetectionResult", string(data)}, " ")
}
