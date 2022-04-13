package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
