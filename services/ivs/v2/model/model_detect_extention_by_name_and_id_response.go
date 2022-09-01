package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectExtentionByNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty" xml:"meta"`

	Result         *IvsExtentionByNameAndIdResponseBodyResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int                                        `json:"-"`
}

func (o DetectExtentionByNameAndIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByNameAndIdResponse", string(data)}, " ")
}
