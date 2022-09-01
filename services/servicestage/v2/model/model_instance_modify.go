package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceModify struct {

	// 应用组件版本号，满足版本语义，如1.0.1。
	Version string `json:"version" xml:"version"`

	FlavorId *FlavorId `json:"flavor_id,omitempty" xml:"flavor_id"`

	// 组件部署件。key为组件component_name，对于Docker多容器场景，key为容器名称。
	Artifacts map[string]interface{} `json:"artifacts,omitempty" xml:"artifacts"`

	// 应用配置，如环境变量。
	Configuration map[string]interface{} `json:"configuration,omitempty" xml:"configuration"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 访问方式列表。
	ExternalAccesses *[]ExternalAccesses `json:"external_accesses,omitempty" xml:"external_accesses"`

	// 部署资源列表。
	ReferResources *[]ReferResourceCreate `json:"refer_resources,omitempty" xml:"refer_resources"`
}

func (o InstanceModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceModify struct{}"
	}

	return strings.Join([]string{"InstanceModify", string(data)}, " ")
}
