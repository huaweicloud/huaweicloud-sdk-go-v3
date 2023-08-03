package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RtcRoomInfoList RTC Animation加入房间信息。
type RtcRoomInfoList struct {

	// RTC应用ID。
	AppId *string `json:"app_id,omitempty"`

	// RTC房间ID。
	RoomId *string `json:"room_id,omitempty"`

	// 接入RTC的用户信息。
	Users *[]RtcUserInfo `json:"users,omitempty"`
}

func (o RtcRoomInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcRoomInfoList struct{}"
	}

	return strings.Join([]string{"RtcRoomInfoList", string(data)}, " ")
}
