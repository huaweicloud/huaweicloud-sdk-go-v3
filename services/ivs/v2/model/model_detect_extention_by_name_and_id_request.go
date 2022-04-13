package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectExtentionByNameAndIdRequest struct {
	Body *IvsExtentionByNameAndIdRequestBody `json:"body,omitempty"`
}

func (o DetectExtentionByNameAndIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectExtentionByNameAndIdRequest struct{}"
	}

	return strings.Join([]string{"DetectExtentionByNameAndIdRequest", string(data)}, " ")
}
