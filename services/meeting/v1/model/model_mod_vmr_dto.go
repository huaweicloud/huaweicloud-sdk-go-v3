package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户修改vmr信息
type ModVmrDto struct {
	// 云会议室名称 maxLength：128 minLength：1

	VmrName *string `json:"vmrName,omitempty"`
	// 来宾密码，“”代表不设置来宾密码 4~16位的数字

	GustPwd *string `json:"gustPwd,omitempty"`
	// 主席密码 4~16位的数字

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
