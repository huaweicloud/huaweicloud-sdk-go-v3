package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOpenSourceEnabledRequest Request Object
type ShowOpenSourceEnabledRequest struct {
}

func (o ShowOpenSourceEnabledRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpenSourceEnabledRequest struct{}"
	}

	return strings.Join([]string{"ShowOpenSourceEnabledRequest", string(data)}, " ")
}
