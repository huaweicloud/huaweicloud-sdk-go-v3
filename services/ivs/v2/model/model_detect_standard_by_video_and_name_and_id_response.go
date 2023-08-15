package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectStandardByVideoAndNameAndIdResponse Response Object
type DetectStandardByVideoAndNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsStandardByVideoAndNameAndIdResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                               `json:"-"`
}

func (o DetectStandardByVideoAndNameAndIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByVideoAndNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByVideoAndNameAndIdResponse", string(data)}, " ")
}
