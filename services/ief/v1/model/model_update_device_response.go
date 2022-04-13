package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDeviceResponse struct {
	Device         *Device `json:"device,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceResponse", string(data)}, " ")
}
