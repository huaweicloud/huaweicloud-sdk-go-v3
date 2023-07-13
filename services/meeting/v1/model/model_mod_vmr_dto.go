package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModVmrDto 用户修改云会议室或者个人会议ID信息。
type ModVmrDto struct {

	// 云会议室名称。
	VmrName *string `json:"vmrName,omitempty"`

	// 来宾密码，设置为空字符串代表不设置来宾密码。 4~16位的数字
	GustPwd *string `json:"gustPwd,omitempty"`

	// 主持人密码。4~16位的数字。
	ChairPwd *string `json:"chairPwd,omitempty"`

	// 是否允许来宾先入会。
	AllowGustFirst *bool `json:"allowGustFirst,omitempty"`

	// 云会议室被使用后是否通知会议室所有者。
	GustFirstNotice *bool `json:"gustFirstNotice,omitempty"`
}

func (o ModVmrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModVmrDto struct{}"
	}

	return strings.Join([]string{"ModVmrDto", string(data)}, " ")
}
