package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByBase64IntlRequest struct {
	Body *FaceDetectBase64Req `json:"body,omitempty"`
}

func (o DetectFaceByBase64IntlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByBase64IntlRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByBase64IntlRequest", string(data)}, " ")
}
