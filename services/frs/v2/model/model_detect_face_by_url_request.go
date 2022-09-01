package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByUrlRequest struct {
	Body *FaceDetectUrlReq `json:"body,omitempty" xml:"body"`
}

func (o DetectFaceByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByUrlRequest", string(data)}, " ")
}
