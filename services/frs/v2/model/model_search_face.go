package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchFace struct {
	BoundingBox *BoundingBox `json:"bounding_box,omitempty"`
	// 人脸搜索时用于被检索的相似度。

	Similarity *float64 `json:"similarity,omitempty"`
	// 用户添加的额外自定义字段。

	ExternalFields *interface{} `json:"external_fields,omitempty"`
	// 人脸所在的外部图片ID。

	ExternalImageId *string `json:"external_image_id,omitempty"`
	// 人脸ID，由系统内部生成的唯一ID。

	FaceId *string `json:"face_id,omitempty"`
}

func (o SearchFace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFace struct{}"
	}

	return strings.Join([]string{"SearchFace", string(data)}, " ")
}
