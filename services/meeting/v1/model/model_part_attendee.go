package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部分与会者信息。
type PartAttendee struct {

	// 与会者名称。
	Name *string `json:"name,omitempty"`

	// 号码。SIP号码或者手机号码。
	Phone *string `json:"phone,omitempty"`

	// 预留字段，取值类型同参数“phone”。
	Phone2 *string `json:"phone2,omitempty"`

	// 预留字段，取值类型同参数“phone”。
	Phone3 *string `json:"phone3,omitempty"`

	// 终端类型，类型枚举如下： * normal：软终端 * terminal：硬终端 * outside：外部与会人 * mobile：用户手机号码 * ideahub：ideahub * board: 电子白板（SmartRooms）。含Maxhub、海信大屏、IdeaHub B2hwvision：华为智慧屏TV
	Type *string `json:"type,omitempty"`

	// 用户入会时是否需要自动静音 。默认不静音。 * 0： 不需要静音 * 1： 需要静音
	Role *int32 `json:"role,omitempty"`

	// 用户入会时是否需要自动静音。默认不静音。 * 0: 不需要静音。 * 1: 需要静音。
	IsMute *int32 `json:"isMute,omitempty"`
}

func (o PartAttendee) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartAttendee struct{}"
	}

	return strings.Join([]string{"PartAttendee", string(data)}, " ")
}
