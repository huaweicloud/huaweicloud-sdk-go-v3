package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveByBase64IntlRequest struct {
	Body *LiveDetectBase64Req `json:"body,omitempty"`
}

func (o DetectLiveByBase64IntlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByBase64IntlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByBase64IntlRequest", string(data)}, " ")
}
