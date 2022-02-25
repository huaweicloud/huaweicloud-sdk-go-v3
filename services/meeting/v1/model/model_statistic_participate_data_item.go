package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会议与会统计数据的单个时间点数据。
type StatisticParticipateDataItem struct {
	// 日期/月份。

	Time *string `json:"time,omitempty"`
	// 与会用户名称。 category = user_participate_info时有效。

	ConfUserName *string `json:"confUserName,omitempty"`
	// 与会用户账户。 category = user_participate_info时有效。

	ConfUserAccount *string `json:"confUserAccount,omitempty"`
	// 与会用户所属部门。 category = user_participate_info时有效。

	ConfUserDeptName *string `json:"confUserDeptName,omitempty"`
	// 用户与会数。 category = user_participate_info时有效。

	ConfUserCount *string `json:"confUserCount,omitempty"`
	// 用户与会时长(秒)。 category = user_participate_info时有效。

	ConfUserDuration *string `json:"confUserDuration,omitempty"`
	// 与会硬件终端名称。 category = hard_terminal_participate_info时有效。

	ConfHardTerminalName *string `json:"confHardTerminalName,omitempty"`
	// 与会硬件终端型号。 category = hard_terminal_participate_info时有效。

	ConfHardTerminalModel *string `json:"confHardTerminalModel,omitempty"`
	// 与会硬件终端的用户ID。 category = hard_terminal_participate_info时有效。

	ConfHardTerminalUserId *string `json:"confHardTerminalUserId,omitempty"`
	// 硬件终端与会数。 category = hard_terminal_participate_info时有效。

	ConfHardTerminalCount *string `json:"confHardTerminalCount,omitempty"`
	// 硬件终端与会时长(秒)。 category = hard_terminal_participate_info时有效。

	ConfHardTerminalDuration *string `json:"confHardTerminalDuration,omitempty"`
	// 与会设备类型。 category = participant_type_info时有效。

	DeviceType *string `json:"deviceType,omitempty"`
	// 与会设备版本。 category = participant_type_info时有效。

	DeviceVersion *string `json:"deviceVersion,omitempty"`
	// 设备与会数。 category = participant_type_info时有效。

	DeviceAttendanceCount *string `json:"deviceAttendanceCount,omitempty"`
}

func (o StatisticParticipateDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticParticipateDataItem struct{}"
	}

	return strings.Join([]string{"StatisticParticipateDataItem", string(data)}, " ")
}
