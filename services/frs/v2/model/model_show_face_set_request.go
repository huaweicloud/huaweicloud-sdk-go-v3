package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFaceSetRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
}

func (o ShowFaceSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFaceSetRequest struct{}"
	}

	return strings.Join([]string{"ShowFaceSetRequest", string(data)}, " ")
}
