package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulHostAppsResponseInfoPaths struct {

	// **参数解释**: 存在漏洞的软件路径 **取值范围**: 字符长度1-512位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符长度1-64位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 漏洞状态 **取值范围**: - vul_status_unfix：未处理 - vul_status_ignored：已忽略 - vul_status_verified：验证中 - vul_status_fixing：修复中 - vul_status_fixed：修复成功 - vul_status_reboot：修复成功待重启 - vul_status_failed：修复失败 - vul_status_fix_after_reboot：请重启主机再次修复
	Status *string `json:"status,omitempty"`
}

func (o VulHostAppsResponseInfoPaths) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostAppsResponseInfoPaths struct{}"
	}

	return strings.Join([]string{"VulHostAppsResponseInfoPaths", string(data)}, " ")
}
