package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceResponse Response Object
type ShowDeviceResponse struct {
	Device         *Device `json:"device,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceResponse", string(data)}, " ")
}
