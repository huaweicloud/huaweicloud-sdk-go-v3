package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例参数。
type InstanceListView struct {

	// 应用组件实例ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用ID。
	ApplicationId *string `json:"application_id,omitempty" xml:"application_id"`

	// 应用名称。
	ApplicationName *string `json:"application_name,omitempty" xml:"application_name"`

	// 组件ID。
	ComponentId *string `json:"component_id,omitempty" xml:"component_id"`

	// 组件名称。
	ComponentName *string `json:"component_name,omitempty" xml:"component_name"`

	// 应用组件实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用组件环境ID。
	EnvironmentId *string `json:"environment_id,omitempty" xml:"environment_id"`

	// 环境名称。
	EnvironmentName *string `json:"environment_name,omitempty" xml:"environment_name"`

	// 运行平台类型。 应用可以在不同的平台上运行，可选用的平台的类型有以下几种：cce、vmapp。
	PlatformType *string `json:"platform_type,omitempty" xml:"platform_type"`

	// 应用组件版本号。
	Version *string `json:"version,omitempty" xml:"version"`

	// 访问方式。
	ExternalAccesses *[]ExternalAccesses `json:"external_accesses,omitempty" xml:"external_accesses"`

	// 组件部署件。key为组件component_name，对于Docker多容器场景，key为容器名称。
	Artifacts map[string]interface{} `json:"artifacts,omitempty" xml:"artifacts"`

	// 创建人。
	Creator *string `json:"creator,omitempty" xml:"creator"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	StatusDetail *InstanceStatusView `json:"status_detail,omitempty" xml:"status_detail"`
}

func (o InstanceListView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceListView struct{}"
	}

	return strings.Join([]string{"InstanceListView", string(data)}, " ")
}
