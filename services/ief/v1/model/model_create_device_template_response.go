package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDeviceTemplateResponse struct {
	DeviceTemplate *EdgemgrDevice `json:"device_template,omitempty" xml:"device_template"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateDeviceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateDeviceTemplateResponse", string(data)}, " ")
}
