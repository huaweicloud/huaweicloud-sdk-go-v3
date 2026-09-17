package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddonTemplatesRequest Request Object
type ListAddonTemplatesRequest struct {

	// **参数解释**： 指定的插件名称或插件别名，不填写则查询列表。 **约束限制**： 不涉及 **取值范围**： 以字母或数字开头和结尾，支持字母、数字、连字符(-)、下划线(_)和点号(.)，长度范围2~30位。 **默认取值**： 不涉及
	AddonTemplateName *string `json:"addon_template_name,omitempty"`

	// 含义：可接受的最低升级版本  属性：隐藏参数
	BaseUpdateAddonVersion *string `json:"base_update_addon_version,omitempty"`

	// 含义：查询的集群  属性：隐藏参数
	ClusterId *string `json:"cluster_id,omitempty"`

	// 含义：是否获取最新插件  属性：隐藏参数
	Newest *string `json:"newest,omitempty"`

	// 含义：筛选的插件版本  属性：隐藏参数
	Version *string `json:"version,omitempty"`
}

func (o ListAddonTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAddonTemplatesRequest", string(data)}, " ")
}
