package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteOtpDevicesResponse Response Object
type BatchDeleteOtpDevicesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteOtpDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOtpDevicesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteOtpDevicesResponse", string(data)}, " ")
}
