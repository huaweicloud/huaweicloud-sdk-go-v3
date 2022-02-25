package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeviceStatusResponse struct {
	Body           *[]UserStatusDto `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDeviceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceStatusResponse", string(data)}, " ")
}
