package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectExtentionByIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty" xml:"meta"`

	Result         *IvsExtentionByIdCardImageResponseBodyResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                                          `json:"-"`
}

func (o DetectExtentionByIdCardImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByIdCardImageResponse", string(data)}, " ")
}
