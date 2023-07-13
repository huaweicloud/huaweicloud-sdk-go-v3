package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceTwinResponse Response Object
type UpdateDeviceTwinResponse struct {
	PropertyVisitors *ValueInPropertyVisitors `json:"property_visitors,omitempty"`

	Twin           *ValueInTwinResponse `json:"twin,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateDeviceTwinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTwinResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTwinResponse", string(data)}, " ")
}
