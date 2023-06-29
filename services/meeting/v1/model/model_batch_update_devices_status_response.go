package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateDevicesStatusResponse Response Object
type BatchUpdateDevicesStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateDevicesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDevicesStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateDevicesStatusResponse", string(data)}, " ")
}
