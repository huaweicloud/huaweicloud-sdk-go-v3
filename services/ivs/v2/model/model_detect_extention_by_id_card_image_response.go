package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectExtentionByIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsExtentionByIdCardImageResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o DetectExtentionByIdCardImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByIdCardImageResponse", string(data)}, " ")
}
