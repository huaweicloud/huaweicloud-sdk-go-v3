package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveByBase64Request struct {
	Body *LiveDetectBase64Req `json:"body,omitempty"`
}

func (o DetectLiveByBase64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByBase64Request struct{}"
	}

	return strings.Join([]string{"DetectLiveByBase64Request", string(data)}, " ")
}
