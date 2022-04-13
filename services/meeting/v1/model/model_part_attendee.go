package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部分与会者信息
type PartAttendee struct {
	// 与会者名称或昵称。长度限制为96个字符。

	Name *string `json:"name,omitempty"`
	// 电话号码(可支持SIP、TEL号码格式)。最大不超过127个字符。 当type为telepresence时，且设备为三屏智真，则该字段填写中屏号码。

	Phone *string `json:"phone,omitempty"`
	// 取值类型同参数phone。（预留字段） 当type为telepresence时，且设备为三屏智真，则该字段填写左屏号码。

	Phone2 *string `json:"phone2,omitempty"`
	// 取值类型同参数phone。（预留字段） 当type为telepresence时，且设备为三屏智真，则该字段填写右屏号码。

	Phone3 *string `json:"phone3,omitempty"`
	// 默认值由会议AS定义，号码类型枚举如下： - normal: 软终端。 - telepresence: 智真。单屏、三屏智真均属此类。（预留字段） - terminal: 会议室或硬终端。 - outside: 外部与会人。 - mobile: 用户手机号码。 - telephone: 用户固定电话。（预留字段） - ideahub: ideahub。

	Type *string `json:"type,omitempty"`
}

func (o PartAttendee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartAttendee struct{}"
	}

	return strings.Join([]string{"PartAttendee", string(data)}, " ")
}
