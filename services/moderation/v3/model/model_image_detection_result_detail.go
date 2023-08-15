package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageDetectionResultDetail 返回结果的详细内容。
type ImageDetectionResultDetail struct {

	// 审核结果是否通过。 block：包含敏感信息，不通过 review：需要人工复检
	Suggestion string `json:"suggestion"`

	// 检测结果的一级标签。 支持category列表如下： terrorism: 暴恐 porn: 色情 image_text: 图文审核
	Category string `json:"category"`

	// 置信度，可选值在0-1之间，值越大，可信度越高。
	Confidence float32 `json:"confidence"`

	FaceLocation *FaceLocationDetail `json:"face_location,omitempty"`

	QrLocation *QrLocationDetail `json:"qr_location,omitempty"`

	// 图片中二维码指向的链接，当请求参数categories中包含image_text时存在。
	QrContent *string `json:"qr_content,omitempty"`

	// image_text场景下命中的文本片段。
	Segments *[]OcrTextDetail `json:"segments,omitempty"`

	// 识别的详细标签。
	Label *string `json:"label,omitempty"`
}

func (o ImageDetectionResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetail", string(data)}, " ")
}
