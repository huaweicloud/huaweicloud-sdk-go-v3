package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByFileRequest struct {
	Body *DetectFaceByFileRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o DetectFaceByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileRequest", string(data)}, " ")
}
