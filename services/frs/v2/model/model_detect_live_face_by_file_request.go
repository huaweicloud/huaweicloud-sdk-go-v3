package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveFaceByFileRequest struct {
	Body *DetectLiveFaceByFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o DetectLiveFaceByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByFileRequest", string(data)}, " ")
}
