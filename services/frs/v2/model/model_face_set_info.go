package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceSetInfo struct {

	// 人脸库当中的人脸数量。
	FaceNumber int32 `json:"face_number"`

	// 用户的自定义字段。
	ExternalFields *interface{} `json:"external_fields"`

	// 人脸库ID，随机生成的包含八个字符的字符串。
	FaceSetId string `json:"face_set_id"`

	// 人脸库名称。
	FaceSetName string `json:"face_set_name"`

	// 创建时间。
	CreateDate string `json:"create_date"`

	// 人脸库最大的容量。
	FaceSetCapacity int32 `json:"face_set_capacity"`
}

func (o FaceSetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceSetInfo struct{}"
	}

	return strings.Join([]string{"FaceSetInfo", string(data)}, " ")
}
