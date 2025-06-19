package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetKeepTimeRequest Request Object
type SetKeepTimeRequest struct {
	Body *SetKeepTimeRequestBody `json:"body,omitempty"`
}

func (o SetKeepTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetKeepTimeRequest struct{}"
	}

	return strings.Join([]string{"SetKeepTimeRequest", string(data)}, " ")
}
