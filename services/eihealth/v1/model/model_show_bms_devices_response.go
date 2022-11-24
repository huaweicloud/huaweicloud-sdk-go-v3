package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBmsDevicesResponse struct {

	// 显卡id列表
	Devices        *[]string `json:"devices,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBmsDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBmsDevicesResponse struct{}"
	}

	return strings.Join([]string{"ShowBmsDevicesResponse", string(data)}, " ")
}
