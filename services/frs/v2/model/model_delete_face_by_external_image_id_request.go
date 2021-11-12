package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteFaceByExternalImageIdRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
	// external_image_id。

	ExternalImageId string `json:"external_image_id"`
}

func (o DeleteFaceByExternalImageIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFaceByExternalImageIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteFaceByExternalImageIdRequest", string(data)}, " ")
}
