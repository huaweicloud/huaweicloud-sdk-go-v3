package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoInstallAppReq 自动安装应用。
type AutoInstallAppReq struct {

	// 安装命令(静默安装命令)。 例: ${FILE_PATH} /S 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
	InstallCommand *string `json:"install_command,omitempty"`

	// 卸载命令(静默卸载命令)。 例: msiexec /uninstall ${FILE_PATH} /quiet。 预定义变量将采用以下值: ${FILE_PATH}: 应用安装包在桌面本地的存储路径。
	UninstallCommand *string `json:"uninstall_command,omitempty"`

	// 指定安装用户。
	Users *[]AccountInfo `json:"users,omitempty"`
}

func (o AutoInstallAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoInstallAppReq struct{}"
	}

	return strings.Join([]string{"AutoInstallAppReq", string(data)}, " ")
}
