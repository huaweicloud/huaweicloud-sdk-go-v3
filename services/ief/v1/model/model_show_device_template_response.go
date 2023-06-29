package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceTemplateResponse Response Object
type ShowDeviceTemplateResponse struct {
	DeviceTemplate *EdgemgrDevice `json:"device_template,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDeviceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceTemplateResponse", string(data)}, " ")
}
