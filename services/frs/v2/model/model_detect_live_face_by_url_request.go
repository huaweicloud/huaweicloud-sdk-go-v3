package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectLiveFaceByUrlRequest struct {
	Body *LiveDetectFaceUrlReq `json:"body,omitempty"`
}

func (o DetectLiveFaceByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByUrlRequest", string(data)}, " ")
}
