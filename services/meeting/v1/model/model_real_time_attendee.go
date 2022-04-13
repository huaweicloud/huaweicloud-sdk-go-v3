package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与会者信息
type RealTimeAttendee struct {
	// 与会者帐号（AAPID用户，代表第三方账号）。

	AccountID *string `json:"accountID,omitempty"`
	// 与会者的用户UUID。

	UserUUID *string `json:"userUUID,omitempty"`
	// 与会者名称或昵称，长度限制为96个字符。

	Name *string `json:"name,omitempty"`
	// 与会者设备的注册号码（可支持SIP、TEL号码格式）。最大不超过127个字符。 设备为三屏智真时的中屏号码。

	Phone *string `json:"phone,omitempty"`
	// 设备为三屏智真时的左屏号码（预留）。

	PhoneLeft *string `json:"phoneLeft,omitempty"`
	// 设备为三屏智真时的右屏号码（预留）。

	PhoneRight *string `json:"phoneRight,omitempty"`
}

func (o RealTimeAttendee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeAttendee struct{}"
	}

	return strings.Join([]string{"RealTimeAttendee", string(data)}, " ")
}
