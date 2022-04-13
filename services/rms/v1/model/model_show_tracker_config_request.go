package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTrackerConfigRequest struct {
}

func (o ShowTrackerConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrackerConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowTrackerConfigRequest", string(data)}, " ")
}
