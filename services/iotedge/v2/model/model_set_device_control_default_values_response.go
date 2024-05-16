package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDeviceControlDefaultValuesResponse Response Object
type SetDeviceControlDefaultValuesResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetDeviceControlDefaultValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDeviceControlDefaultValuesResponse struct{}"
	}

	return strings.Join([]string{"SetDeviceControlDefaultValuesResponse", string(data)}, " ")
}
