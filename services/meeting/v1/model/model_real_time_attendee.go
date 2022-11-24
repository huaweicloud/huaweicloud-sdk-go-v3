package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 被邀请与会者信息。
type RealTimeAttendee struct {

	// 与会者的华为云会议帐号。
	AccountID *string `json:"accountID,omitempty"`

	// 与会者的用户UUID。
	UserUUID *string `json:"userUUID,omitempty"`

	// 与会者名称。
	Name *string `json:"name,omitempty"`

	// 与会者号码。
	Phone *string `json:"phone,omitempty"`

	// 设备为三屏智真时的左屏号码。 > 该参数将废弃，请勿使用。
	PhoneLeft *string `json:"phoneLeft,omitempty"`

	// 设备为三屏智真时的右屏号码。 > 该参数将废弃，请勿使用。
	PhoneRight *string `json:"phoneRight,omitempty"`
}

func (o RealTimeAttendee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeAttendee struct{}"
	}

	return strings.Join([]string{"RealTimeAttendee", string(data)}, " ")
}
