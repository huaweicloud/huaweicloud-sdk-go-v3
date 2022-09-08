package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchFaceByFileRequest struct {

	// 人脸库名称。
	FaceSetName string `json:"face_set_name"`

	Body *SearchFaceByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o SearchFaceByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"SearchFaceByFileRequest", string(data)}, " ")
}
