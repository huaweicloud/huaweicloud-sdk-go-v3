package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceCreate struct {

	// 应用组件实例名称。
	Name string `json:"name"`

	// 环境ID。
	EnvironmentId string `json:"environment_id"`

	FlavorId *FlavorId `json:"flavor_id"`

	// 实例副本数。
	Replica int32 `json:"replica"`

	// 组件部署件。key为组件component_name，对于Docker多容器场景，key为容器名称。
	Artifacts map[string]interface{} `json:"artifacts"`

	// 应用组件版本号，满足版本语义，如1.0.0。。
	Version string `json:"version"`

	Configuration *InstanceConfiguration `json:"configuration,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 访问方式。
	ExternalAccesses *[]ExternalAccessesCreate `json:"external_accesses,omitempty"`

	// 部署资源。
	ReferResources []ReferResourceCreate `json:"refer_resources"`
}

func (o InstanceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceCreate struct{}"
	}

	return strings.Join([]string{"InstanceCreate", string(data)}, " ")
}
