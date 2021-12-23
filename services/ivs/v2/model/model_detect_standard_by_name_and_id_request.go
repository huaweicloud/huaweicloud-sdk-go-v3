package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetectStandardByNameAndIdRequest struct {
	Body *IvsStandardByNameAndIdRequestBody `json:"body,omitempty"`
}

func (o DetectStandardByNameAndIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectStandardByNameAndIdRequest struct{}"
	}

	return strings.Join([]string{"DetectStandardByNameAndIdRequest", string(data)}, " ")
}
