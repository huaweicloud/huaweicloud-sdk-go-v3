package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLandingZoneVersionResponse Response Object
type ListLandingZoneVersionResponse struct {

	// Landing Zone版本列表。
	Body           *[]LandingZoneVersionDetail `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListLandingZoneVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLandingZoneVersionResponse struct{}"
	}

	return strings.Join([]string{"ListLandingZoneVersionResponse", string(data)}, " ")
}
