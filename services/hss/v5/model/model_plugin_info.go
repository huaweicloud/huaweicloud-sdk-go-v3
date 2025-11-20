package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginInfo 插件信息
type PluginInfo struct {

	// **参数解释**： 插件编码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**： 插件名称 **取值范围**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件的描述信息 **取值范围**: 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 插件标签信息 **取值范围**: 不涉及
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释**： 已安装插件的主机数量 **取值范围**: 不涉及
	InstalledAttachmentNum *int32 `json:"installed_attachment_num,omitempty"`

	// **参数解释**： 未安装插件的主机数量，包括插件状态为未安装、安装中与安装失败的主机 **取值范围**: 不涉及
	UninstallAttachmentNum *int32 `json:"uninstall_attachment_num,omitempty"`

	// **参数解释**： 此种类型的插件包中，运行插件所需单核CPU(0-100%)的最大值 **取值范围**: 不涉及
	MaxCpuLimit *int32 `json:"max_cpu_limit,omitempty"`

	// **参数解释**： 此种类型的插件包中，运行插件所需内存(MB)的最大值 **取值范围**: 不涉及
	MaxMemoryLimit *int32 `json:"max_memory_limit,omitempty"`

	// **参数解释**： 此种类型的插件包中，插件大小(MB)的最大值 **取值范围**: 不涉及
	MaxSize *string `json:"max_size,omitempty"`

	// **参数解释**： 插件展示模式 **取值范围**: - 0：插件所有操作功能均支持 - 1：不支持插件的安装、卸载操作 - 2：插件所有操作功能均不支持
	DisplayMode *int32 `json:"display_mode,omitempty"`
}

func (o PluginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginInfo struct{}"
	}

	return strings.Join([]string{"PluginInfo", string(data)}, " ")
}
