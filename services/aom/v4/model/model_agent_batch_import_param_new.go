package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentBatchImportParamNew struct {

	// 导入待安装UniAgent的机器参数列表。
	AgentImportParamList []AgentImportParamNew `json:"agent_import_param_list"`

	// 代理区域ID： - 直连接入填0。 - 代理接入填实际代理区域ID。
	ProxyRegionId int32 `json:"proxy_region_id"`

	// 安装机（代理机）的agent ID。
	InstallerAgentId string `json:"installer_agent_id"`

	// 是否需要安装ICAgent插件： - true：安装ICAgent插件。默认安装最新版本的ICAgent插件。 - false：不安装ICAgent插件。
	IcagentInstallFlag *bool `json:"icagent_install_flag,omitempty"`

	PluginInstallBaseParam *PluginInstallBasicParam `json:"plugin_install_base_param,omitempty"`

	// 待安装的UniAgent版本号。
	Version string `json:"version"`

	// 是否公网接入： - true：公网接入设置。 - false：代理接入设置。
	PublicNetFlag bool `json:"public_net_flag"`
}

func (o AgentBatchImportParamNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentBatchImportParamNew struct{}"
	}

	return strings.Join([]string{"AgentBatchImportParamNew", string(data)}, " ")
}
