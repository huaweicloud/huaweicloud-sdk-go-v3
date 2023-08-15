package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatisticUserDataItem 会议用户统计数据的单个时间点数据。
type StatisticUserDataItem struct {

	// 日期/月份。
	Time *string `json:"time,omitempty"`

	// 登录用户数。 category = user_login_info时有效。
	UserLoginCount *string `json:"userLoginCount,omitempty"`

	// PC端登录用户数。 category = user_login_info时有效。
	UserPCLoginCount *string `json:"userPCLoginCount,omitempty"`

	// 移动端登录用户数。 category = user_login_info时有效。
	UserMobileLoginCount *string `json:"userMobileLoginCount,omitempty"`

	// 激活用户数。 category = user_activate_info时有效。
	UserActivatedCount *string `json:"userActivatedCount,omitempty"`

	// 用户登录设备名称。 category = user_login_device_info时有效。
	UserLoginDevicesName *string `json:"userLoginDevicesName,omitempty"`

	// 用户登录设备数。 category = user_login_device_info时有效。
	UserLoginDevicesCount *string `json:"userLoginDevicesCount,omitempty"`
}

func (o StatisticUserDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticUserDataItem struct{}"
	}

	return strings.Join([]string{"StatisticUserDataItem", string(data)}, " ")
}
