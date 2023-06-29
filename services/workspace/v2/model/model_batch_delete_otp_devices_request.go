package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteOtpDevicesRequest Request Object
type BatchDeleteOtpDevicesRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`

	Body *DelOtpDevicesReq `json:"body,omitempty"`
}

func (o BatchDeleteOtpDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOtpDevicesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteOtpDevicesRequest", string(data)}, " ")
}
