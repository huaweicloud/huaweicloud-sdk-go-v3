package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginVersionResponse Response Object
type ShowPluginVersionResponse struct {

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 操作人
	OpUser *string `json:"op_user,omitempty"`

	// 操作时间
	OpTime *string `json:"op_time,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 版本说明
	VersionDescription *string `json:"version_description,omitempty"`

	// 版本属性
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// 组合插件类型
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// 插件属性
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// 输入信息
	InputInfo *[]PluginPartQueryVoListAgentPluginInputVoData `json:"input_info,omitempty"`

	// 执行信息
	PluginExecution *interface{} `json:"plugin_execution,omitempty"`

	// 运行属性
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowPluginVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginVersionResponse", string(data)}, " ")
}
