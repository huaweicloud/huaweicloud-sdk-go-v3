package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveByFileRequest struct {
	Body *DetectLiveByFileRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o DetectLiveByFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByFileRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByFileRequest", string(data)}, " ")
}
