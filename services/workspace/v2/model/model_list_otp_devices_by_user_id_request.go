package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOtpDevicesByUserIdRequest Request Object
type ListOtpDevicesByUserIdRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`
}

func (o ListOtpDevicesByUserIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOtpDevicesByUserIdRequest struct{}"
	}

	return strings.Join([]string{"ListOtpDevicesByUserIdRequest", string(data)}, " ")
}
