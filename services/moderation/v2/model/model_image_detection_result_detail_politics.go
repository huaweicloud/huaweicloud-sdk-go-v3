package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageDetectionResultDetailPolitics struct {
	// 置信度，取值范围 0-1。

	Confidence *float32 `json:"confidence,omitempty"`
	// 对应的政治人物信息。

	Label *string `json:"label,omitempty"`

	FaceDetail *ImageDetectionResultDetailFaceDetail `json:"face_detail,omitempty"`
}

func (o ImageDetectionResultDetailPolitics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetailPolitics struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetailPolitics", string(data)}, " ")
}
