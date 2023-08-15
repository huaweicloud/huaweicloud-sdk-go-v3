package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceResponse Response Object
type UpdateDeviceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceResponse", string(data)}, " ")
}
