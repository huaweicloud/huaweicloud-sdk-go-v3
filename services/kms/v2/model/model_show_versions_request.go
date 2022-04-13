package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVersionsRequest struct {
}

func (o ShowVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionsRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionsRequest", string(data)}, " ")
}
