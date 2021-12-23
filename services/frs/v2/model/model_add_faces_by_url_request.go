package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddFacesByUrlRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *AddFacesUrlReq `json:"body,omitempty"`
}

func (o AddFacesByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFacesByUrlRequest struct{}"
	}

	return strings.Join([]string{"AddFacesByUrlRequest", string(data)}, " ")
}
