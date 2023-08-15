package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceSetFace struct {
	BoundingBox *BoundingBox `json:"bounding_box"`

	// 用户添加的额外字段。
	ExternalFields *interface{} `json:"external_fields"`

	// 人脸所在的外部图片ID。
	ExternalImageId string `json:"external_image_id"`

	// 人脸ID，由系统内部生成的唯一ID。
	FaceId string `json:"face_id"`
}

func (o FaceSetFace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceSetFace struct{}"
	}

	return strings.Join([]string{"FaceSetFace", string(data)}, " ")
}
