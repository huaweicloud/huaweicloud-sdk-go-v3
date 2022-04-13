package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveByUrlRequest struct {
	Body *LiveDetectUrlReq `json:"body,omitempty"`
}

func (o DetectLiveByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByUrlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByUrlRequest", string(data)}, " ")
}
