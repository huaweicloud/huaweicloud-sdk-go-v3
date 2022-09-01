package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByFileIntlRequest struct {
	Body *DetectFaceByFileIntlRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o DetectFaceByFileIntlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByFileIntlRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileIntlRequest", string(data)}, " ")
}
