package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDeviceResponse struct {
	Device         *Device `json:"device,omitempty" xml:"device"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceResponse struct{}"
	}

	return strings.Join([]string{"CreateDeviceResponse", string(data)}, " ")
}
