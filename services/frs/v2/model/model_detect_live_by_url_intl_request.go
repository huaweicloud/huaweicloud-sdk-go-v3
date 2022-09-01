package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveByUrlIntlRequest struct {
	Body *LiveDetectUrlReq `json:"body,omitempty" xml:"body"`
}

func (o DetectLiveByUrlIntlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByUrlIntlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveByUrlIntlRequest", string(data)}, " ")
}
