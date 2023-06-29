package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableZoneInfoRequest Request Object
type ShowAvailableZoneInfoRequest struct {
}

func (o ShowAvailableZoneInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableZoneInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableZoneInfoRequest", string(data)}, " ")
}
