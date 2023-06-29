package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopUsedHoursInfo 桌面使用时间信息。
type DesktopUsedHoursInfo struct {

	// 桌面Id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 使用的用户。
	DesktopUsername *string `json:"desktop_username,omitempty"`

	// 桌面使用时间列表。
	UsedInfoList *[]DesktopUsedInfo `json:"used_info_list,omitempty"`
}

func (o DesktopUsedHoursInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopUsedHoursInfo struct{}"
	}

	return strings.Join([]string{"DesktopUsedHoursInfo", string(data)}, " ")
}
