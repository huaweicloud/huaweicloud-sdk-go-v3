package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOtpDevicesByUserIdResponse struct {

	// otp设备
	OtpDevices     *[]OtpDevice `json:"otp_devices,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListOtpDevicesByUserIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOtpDevicesByUserIdResponse struct{}"
	}

	return strings.Join([]string{"ListOtpDevicesByUserIdResponse", string(data)}, " ")
}
