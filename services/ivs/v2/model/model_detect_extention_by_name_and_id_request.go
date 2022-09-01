package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectExtentionByNameAndIdRequest struct {
	Body *IvsExtentionByNameAndIdRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DetectExtentionByNameAndIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByNameAndIdRequest struct{}"
	}

	return strings.Join([]string{"DetectExtentionByNameAndIdRequest", string(data)}, " ")
}
