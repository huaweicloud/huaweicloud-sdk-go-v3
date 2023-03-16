package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectStandardByVideoAndIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsStandardByVideoAndIdCardImageResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                                 `json:"-"`
}

func (o DetectStandardByVideoAndIdCardImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByVideoAndIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByVideoAndIdCardImageResponse", string(data)}, " ")
}
