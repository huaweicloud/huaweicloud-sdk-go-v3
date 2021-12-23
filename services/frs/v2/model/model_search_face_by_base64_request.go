package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchFaceByBase64Request struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *FaceSearchBase64Req `json:"body,omitempty"`
}

func (o SearchFaceByBase64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFaceByBase64Request struct{}"
	}

	return strings.Join([]string{"SearchFaceByBase64Request", string(data)}, " ")
}
