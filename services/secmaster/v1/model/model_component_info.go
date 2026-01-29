package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentInfo 插件详细信息
type ComponentInfo struct {

	// 插件ID
	Id *string `json:"id,omitempty"`

	// 插件名称
	Name *string `json:"name,omitempty"`

	// 插件的实现语言
	DevLanguage *string `json:"dev_language,omitempty"`

	// 插件的实现语言版本
	DevLanguageVersion *string `json:"dev_language_version,omitempty"`

	// 插件集ID
	AllianceId *string `json:"alliance_id,omitempty"`

	// 插件集名称
	AllianceName *string `json:"alliance_name,omitempty"`

	// 插件描述
	Description *string `json:"description,omitempty"`

	// 插件图标
	Logo *string `json:"logo,omitempty"`

	// 插件的标签信息
	Label *string `json:"label,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改者ID
	ModifierId *string `json:"modifier_id,omitempty"`

	// 修改者名称
	ModifierName *string `json:"modifier_name,omitempty"`

	// 插件操作历史
	OperationHistory *[]ComponentInfoOperationHistory `json:"operation_history,omitempty"`

	// 插件版本信息，兼容之前java的以插件为粒度的版本
	Versions *[]ComponentVersionInfo `json:"versions,omitempty"`

	// 插件类型，subscribe/custom/system，对应订阅/自定义开发/系统内置
	ComponentType *string `json:"component_type,omitempty"`
}

func (o ComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInfo struct{}"
	}

	return strings.Join([]string{"ComponentInfo", string(data)}, " ")
}
