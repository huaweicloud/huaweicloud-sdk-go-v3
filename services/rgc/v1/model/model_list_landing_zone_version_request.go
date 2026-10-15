package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLandingZoneVersionRequest Request Object
type ListLandingZoneVersionRequest struct {
}

func (o ListLandingZoneVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLandingZoneVersionRequest struct{}"
	}

	return strings.Join([]string{"ListLandingZoneVersionRequest", string(data)}, " ")
}
