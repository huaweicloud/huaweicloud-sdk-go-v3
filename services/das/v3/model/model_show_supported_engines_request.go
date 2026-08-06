package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupportedEnginesRequest Request Object
type ShowSupportedEnginesRequest struct {
}

func (o ShowSupportedEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupportedEnginesRequest struct{}"
	}

	return strings.Join([]string{"ShowSupportedEnginesRequest", string(data)}, " ")
}
