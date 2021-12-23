package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByBase64Request struct {
	Body *FaceDetectBase64Req `json:"body,omitempty"`
}

func (o DetectFaceByBase64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByBase64Request struct{}"
	}

	return strings.Join([]string{"DetectFaceByBase64Request", string(data)}, " ")
}
