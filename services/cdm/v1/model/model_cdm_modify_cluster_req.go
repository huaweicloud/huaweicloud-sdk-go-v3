package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmModifyClusterReq struct {

	// 自动关机。
	AutoOff *bool `json:"autoOff,omitempty"`

	// 定时关机。
	ScheduleBootOff *bool `json:"scheduleBootOff,omitempty"`

	// 定时开机。
	ScheduleBootTime *string `json:"scheduleBootTime,omitempty"`

	// 定时关机时间。
	ScheduleOffTime *string `json:"scheduleOffTime,omitempty"`

	// 消息通知。
	AutoRemind *bool `json:"autoRemind,omitempty"`

	// 手机号码，最多填写20个，以英文逗号分隔。
	PhoneNum *string `json:"phoneNum,omitempty"`

	// 邮箱地址，最多填写20个，以英文逗号分隔。
	Email *string `json:"email,omitempty"`
}

func (o CdmModifyClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmModifyClusterReq struct{}"
	}

	return strings.Join([]string{"CdmModifyClusterReq", string(data)}, " ")
}
