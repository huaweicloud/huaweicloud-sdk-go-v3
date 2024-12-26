package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentInstallScriptResponseInfo agent安装脚本
type AgentInstallScriptResponseInfo struct {

	// 包类型
	PackageType *string `json:"package_type,omitempty"`

	// 命令
	Cmd *string `json:"cmd,omitempty"`

	// 包下载url
	PackageDownloadUrl *string `json:"package_download_url,omitempty"`
}

func (o AgentInstallScriptResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentInstallScriptResponseInfo struct{}"
	}

	return strings.Join([]string{"AgentInstallScriptResponseInfo", string(data)}, " ")
}
