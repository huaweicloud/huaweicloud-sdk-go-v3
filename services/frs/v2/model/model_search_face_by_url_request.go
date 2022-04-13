package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchFaceByUrlRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *FaceSearchUrlReq `json:"body,omitempty"`
}

func (o SearchFaceByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"SearchFaceByUrlRequest", string(data)}, " ")
}
