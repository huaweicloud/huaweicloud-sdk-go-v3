package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddonTemplatesRequest Request Object
type ListAddonTemplatesRequest struct {

	// 插件包版本号
	Version *string `json:"version,omitempty"`

	// 是否获取集群内创建的最新状态
	Newest *string `json:"newest,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 插件模板名称
	AddonTemplateName *string `json:"addon_template_name,omitempty"`

	// 插件的最低版本
	BaseUpdateAddonVersion *string `json:"base_update_addon_version,omitempty"`
}

func (o ListAddonTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAddonTemplatesRequest", string(data)}, " ")
}
