package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonInstanceSpec 插件实例信息
type AddonInstanceSpec struct {

	// cluster ID信息
	ClusterID *string `json:"clusterID,omitempty"`

	// 插件版本信息
	Version *string `json:"version,omitempty"`

	// 插件模板名称
	AddonTemplateName *string `json:"addonTemplateName,omitempty"`

	// 插件模板类型
	AddonTemplateType *string `json:"addonTemplateType,omitempty"`

	// 插件模板Logo
	AddonTemplateLogo *string `json:"addonTemplateLogo,omitempty"`

	// 插件模板标签
	AddonTemplateLabels *[]string `json:"addonTemplateLabels,omitempty"`

	// 信息说明
	Description *string `json:"description,omitempty"`

	// 插件实例的配置参数
	Values map[string]interface{} `json:"values,omitempty"`

	Parameters *ReleaseParams `json:"parameters,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o AddonInstanceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInstanceSpec struct{}"
	}

	return strings.Join([]string{"AddonInstanceSpec", string(data)}, " ")
}
