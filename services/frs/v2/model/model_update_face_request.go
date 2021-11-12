package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFaceRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *UpdateFaceReq `json:"body,omitempty"`
}

func (o UpdateFaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateFaceRequest", string(data)}, " ")
}
