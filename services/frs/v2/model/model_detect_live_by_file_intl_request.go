package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveByFileIntlRequest struct {
	Body *DetectLiveByFileIntlRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o DetectLiveByFileIntlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByFileIntlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByFileIntlRequest", string(data)}, " ")
}
