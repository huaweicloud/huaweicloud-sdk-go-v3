package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceResponse Response Object
type DeleteDeviceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceResponse", string(data)}, " ")
}
