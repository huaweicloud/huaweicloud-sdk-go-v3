package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeWebImageRequest struct {
	Body *WebImageRequestBody `json:"body,omitempty"`
}

func (o RecognizeWebImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeWebImageRequest struct{}"
	}

	return strings.Join([]string{"RecognizeWebImageRequest", string(data)}, " ")
}
