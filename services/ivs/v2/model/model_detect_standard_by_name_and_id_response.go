package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectStandardByNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsStandardByNameAndIdResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o DetectStandardByNameAndIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByNameAndIdResponse", string(data)}, " ")
}
