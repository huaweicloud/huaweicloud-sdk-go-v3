package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchAppWhitelistPolicyLearnStatusRequestInfo struct {

	// 操作状态
	SwitchStatus *int32 `json:"switch_status,omitempty"`

	// 主机列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o SwitchAppWhitelistPolicyLearnStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAppWhitelistPolicyLearnStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchAppWhitelistPolicyLearnStatusRequestInfo", string(data)}, " ")
}
