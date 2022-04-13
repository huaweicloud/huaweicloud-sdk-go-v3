package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectExtentionByIdCardImageRequest struct {
	Body *IvsExtentionByIdCardImageRequestBody `json:"body,omitempty"`
}

func (o DetectExtentionByIdCardImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByIdCardImageRequest struct{}"
	}

	return strings.Join([]string{"DetectExtentionByIdCardImageRequest", string(data)}, " ")
}
