package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCloudPhoneDetailRequest struct {

	// 云手机id。
	PhoneId string `json:"phone_id"`
}

func (o ShowCloudPhoneDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneDetailRequest", string(data)}, " ")
}
