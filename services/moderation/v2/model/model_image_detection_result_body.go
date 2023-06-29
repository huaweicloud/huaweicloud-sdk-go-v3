package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageDetectionResultBody
type ImageDetectionResultBody struct {
	Detail *ImageDetectionResultDetail `json:"detail,omitempty"`

	// 检测结果是否通过。 - block：包含敏感信息，不通过 - pass：不包含敏感信息，通过 - review：需要人工复检 > 当同时检测多个场景时，suggestion的值以最可能包含敏感信息的场景为准。即任一场景出现了block则总的suggestion为block，所有场景都pass时suggestion为pass，这两种情况之外则一定有场景需要review，此时suggestion为review。
	Suggestion *string `json:"suggestion,omitempty"`

	// 具体每个场景的检测结果。  block：包含敏感信息，不通过  pass：不包含敏感信息，通过  review：需要人工复检
	CategorySuggestions map[string]string `json:"category_suggestions,omitempty"`

	// 文本结果
	OcrText *string `json:"ocr_text,omitempty"`
}

func (o ImageDetectionResultBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultBody struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultBody", string(data)}, " ")
}
