package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByUrlIntlRequest struct {
	Body *FaceDetectUrlReq `json:"body,omitempty"`
}

func (o DetectFaceByUrlIntlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByUrlIntlRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByUrlIntlRequest", string(data)}, " ")
}
