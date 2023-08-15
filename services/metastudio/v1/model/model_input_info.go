package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InputInfo 输入信息。
type InputInfo struct {
	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`
}

func (o InputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputInfo struct{}"
	}

	return strings.Join([]string{"InputInfo", string(data)}, " ")
}
