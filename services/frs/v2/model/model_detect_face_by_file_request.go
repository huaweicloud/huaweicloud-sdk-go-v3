package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectFaceByFileRequest struct {
	Body *DetectFaceByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o DetectFaceByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileRequest", string(data)}, " ")
}
