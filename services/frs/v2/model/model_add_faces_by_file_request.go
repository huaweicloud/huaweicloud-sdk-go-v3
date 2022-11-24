package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddFacesByFileRequest struct {

	// 人脸库名称。
	FaceSetName string `json:"face_set_name"`

	Body *AddFacesByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o AddFacesByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFacesByFileRequest struct{}"
	}

	return strings.Join([]string{"AddFacesByFileRequest", string(data)}, " ")
}
