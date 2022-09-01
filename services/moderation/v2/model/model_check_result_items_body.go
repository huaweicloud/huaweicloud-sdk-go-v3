package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CheckResultItemsBody struct {

	// 图片的URL路径。
	Url *string `json:"url,omitempty" xml:"url"`

	// 检测结果是否通过。 - block：包含敏感信息，不通过 - pass：不包含敏感信息，通过 - review：需要人工复检 > 当同时检测多个场景时，suggestion的值以最可能包含敏感信息的场景为准。即任一场景出现了block则总的suggestion为block，所有场景都pass时suggestion为pass，这两种情况之外则一定有场景需要review，此时suggestion为review。
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion"`

	Detail *ImageDetectionResultDetail `json:"detail,omitempty" xml:"detail"`

	// 具体每个场景的检测结果。  block：包含敏感信息，不通过  pass：不包含敏感信息，通过  review：需要人工复检
	CategorySuggestions map[string]string `json:"category_suggestions,omitempty" xml:"category_suggestions"`

	// ocr识别结果。
	OcrText *string `json:"ocr_text,omitempty" xml:"ocr_text"`

	// 图像审核失败时错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 图像审核失败时错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o CheckResultItemsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResultItemsBody struct{}"
	}

	return strings.Join([]string{"CheckResultItemsBody", string(data)}, " ")
}
