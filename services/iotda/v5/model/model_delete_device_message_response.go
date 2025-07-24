package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceMessageResponse Response Object
type DeleteDeviceMessageResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeviceMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceMessageResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceMessageResponse", string(data)}, " ")
}
