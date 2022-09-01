package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectStandardByIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty" xml:"meta"`

	Result         *IvsStandardByIdCardImageResponseBodyResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                                         `json:"-"`
}

func (o DetectStandardByIdCardImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByIdCardImageResponse", string(data)}, " ")
}
