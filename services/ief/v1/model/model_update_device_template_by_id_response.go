package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDeviceTemplateByIdResponse struct {
	DeviceTemplate *EdgemgrDevice `json:"device_template,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateDeviceTemplateByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTemplateByIdResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTemplateByIdResponse", string(data)}, " ")
}
