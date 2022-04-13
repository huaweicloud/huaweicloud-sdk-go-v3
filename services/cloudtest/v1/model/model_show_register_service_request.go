package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRegisterServiceRequest struct {
}

func (o ShowRegisterServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRegisterServiceRequest struct{}"
	}

	return strings.Join([]string{"ShowRegisterServiceRequest", string(data)}, " ")
}
