package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginDetailInfo 插件详细信息
type PluginDetailInfo struct {

	// **参数解释**： 插件名称 **取值范围**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件id **取值范围**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 插件版本 **取值范围**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件支持的最低Agent版本 **取值范围**: 不涉及
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释**： 插件架构 **取值范围**: - x86_64：X86架构 - arm：ARM架构
	Arch *string `json:"arch,omitempty"`

	// **参数解释**： 插件支持的操作系统类型 **取值范围**: - Linux：Linux操作系统 - Windows：Windows操作系统
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 插件版本信息描述 **取值范围**: 不涉及
	VersionDescription *string `json:"version_description,omitempty"`

	// **参数解释**： 插件安装包大小(MB) **取值范围**: 不涉及
	Size *string `json:"size,omitempty"`

	// **参数解释**： 运行插件所需单核CPU(0-100%) **取值范围**: 不涉及
	CpuLimit *int32 `json:"cpu_limit,omitempty"`

	// **参数解释**： 运行插件所需内存(MB) **取值范围**: 不涉及
	MemoryLimit *int32 `json:"memory_limit,omitempty"`

	// **参数解释**： 插件更新时间 **取值范围**: 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o PluginDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDetailInfo struct{}"
	}

	return strings.Join([]string{"PluginDetailInfo", string(data)}, " ")
}
