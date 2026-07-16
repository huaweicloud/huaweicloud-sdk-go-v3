package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSpec spec是集合类的元素类型，内容为插件实例具体信息
type InstanceSpec struct {

	// 集群id
	ClusterID string `json:"clusterID"`

	// 插件模板版本号，如1.0.0
	Version string `json:"version"`

	// 插件模板名称，如coredns
	AddonTemplateName string `json:"addonTemplateName"`

	// **参数解释：** 插件模板类型 **约束限制：** 不涉及 **取值范围：** - helm: 表示使用Helm包进行部署的模板类型 - static: 表示静态模板类型  **默认取值：** 不涉及
	AddonTemplateType string `json:"addonTemplateType"`

	// 插件模板logo图片的地址
	AddonTemplateLogo *string `json:"addonTemplateLogo,omitempty"`

	// 插件模板所属类型
	AddonTemplateLabels *[]string `json:"addonTemplateLabels,omitempty"`

	// 插件模板描述
	Description string `json:"description"`

	// 插件模板安装参数（各插件不同），请根据具体插件模板信息填写安装参数。
	Values map[string]interface{} `json:"values"`
}

func (o InstanceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSpec struct{}"
	}

	return strings.Join([]string{"InstanceSpec", string(data)}, " ")
}
