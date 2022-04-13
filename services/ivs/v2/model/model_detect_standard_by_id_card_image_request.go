package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectStandardByIdCardImageRequest struct {
	Body *IvsStandardByIdCardImageRequestBody `json:"body,omitempty"`
}

func (o DetectStandardByIdCardImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByIdCardImageRequest struct{}"
	}

	return strings.Join([]string{"DetectStandardByIdCardImageRequest", string(data)}, " ")
}
