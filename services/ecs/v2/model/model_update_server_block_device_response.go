package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateServerBlockDeviceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerBlockDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerBlockDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerBlockDeviceResponse", string(data)}, " ")
}
