package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAppWhitelistPolicyHostRequestInfo 主机开启/关闭白名单策略防护
type SwitchAppWhitelistPolicyHostRequestInfo struct {

	// 是否开启/关闭白名单策略防护
	EnableAppWhitelist bool `json:"enable_app_whitelist"`

	// 策略关联主机列表
	PolicyInfoList []AppPolicyInfoList `json:"policy_info_list"`
}

func (o SwitchAppWhitelistPolicyHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAppWhitelistPolicyHostRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchAppWhitelistPolicyHostRequestInfo", string(data)}, " ")
}
