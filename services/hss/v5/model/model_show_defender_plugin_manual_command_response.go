package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefenderPluginManualCommandResponse Response Object
type ShowDefenderPluginManualCommandResponse struct {

	// **参数解释**： 操作类型，回显请求的opt值 **取值范围**： - install：安装 - uninstall：卸载
	Opt *string `json:"opt,omitempty"`

	// **参数解释**： 完整的手动操作命令，可直接复制到终端执行 **取值范围**： 字符长度为1-4096位
	ClawPluginCommand *string `json:"claw_plugin_command,omitempty"`

	// **参数解释**： 脚本下载地址 **取值范围**： 字符长度为1-512位
	ScriptUrl *string `json:"script_url,omitempty"`

	// **参数解释**： 安装凭证 **取值范围**： 字符长度为1-256位
	InstallKey *string `json:"install_key,omitempty"`

	// **参数解释**： 插件包下载地址 **取值范围**： 字符长度为1-512位
	PluginPkgUrl *string `json:"plugin_pkg_url,omitempty"`

	// **参数解释**： AI Defender服务端地址 **取值范围**： 字符长度为1-512位
	Endpoint *string `json:"endpoint,omitempty"`

	// **参数解释**： 安装凭证过期时间，毫秒时间戳，过期后需重新获取 **取值范围**： 最小值0，最大值9223372036854775807
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释**： 查询安装结果命令 **取值范围**： \"openclaw plugins info AgentDR_OpenClaw\"
	QueryResultCmd *string `json:"query_result_cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDefenderPluginManualCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefenderPluginManualCommandResponse struct{}"
	}

	return strings.Join([]string{"ShowDefenderPluginManualCommandResponse", string(data)}, " ")
}
