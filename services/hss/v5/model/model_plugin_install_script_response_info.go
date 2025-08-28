package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginInstallScriptResponseInfo docker插件安装脚本详情
type PluginInstallScriptResponseInfo struct {

	// **参数解释**： 安装包类型 **取值范围**： 字符长度0-128位
	PackageType *string `json:"package_type,omitempty"`

	// **参数解释**： 命令cmd **取值范围**： 字符长度1-256位
	Cmd *string `json:"cmd,omitempty"`

	// **参数解释**： 包下载地址 **取值范围**： 字符长度1-256位
	PackageDownloadUrl *string `json:"package_download_url,omitempty"`
}

func (o PluginInstallScriptResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginInstallScriptResponseInfo struct{}"
	}

	return strings.Join([]string{"PluginInstallScriptResponseInfo", string(data)}, " ")
}
