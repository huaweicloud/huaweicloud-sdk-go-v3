package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProcessRequest struct {
}

func (o ShowProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProcessRequest struct{}"
	}

	return strings.Join([]string{"ShowProcessRequest", string(data)}, " ")
}
