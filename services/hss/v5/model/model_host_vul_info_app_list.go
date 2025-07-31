package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostVulInfoAppList struct {

	// **参数解释**: 软件名称 **取值范围**: 字符范围0-256位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符范围0-256位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 修复漏洞软件需要升级到的版本 **取值范围**: 字符范围0-256位
	UpgradeVersion *string `json:"upgrade_version,omitempty"`

	// **参数解释**: 应用软件的路径（只有应用漏洞有该字段） **取值范围**: 字符范围1-512位
	AppPath *string `json:"app_path,omitempty"`
}

func (o HostVulInfoAppList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulInfoAppList struct{}"
	}

	return strings.Join([]string{"HostVulInfoAppList", string(data)}, " ")
}
