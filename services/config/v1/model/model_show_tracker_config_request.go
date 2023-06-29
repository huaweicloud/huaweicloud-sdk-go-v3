package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrackerConfigRequest Request Object
type ShowTrackerConfigRequest struct {
}

func (o ShowTrackerConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrackerConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowTrackerConfigRequest", string(data)}, " ")
}
