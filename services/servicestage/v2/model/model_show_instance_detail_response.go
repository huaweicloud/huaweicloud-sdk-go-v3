package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceDetailResponse struct {

	// 应用组件实例ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用组件实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 实例描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 应用组件环境ID。
	EnvironmentId *string `json:"environment_id,omitempty" xml:"environment_id"`

	PlatformType *InstancePlatformType `json:"platform_type,omitempty" xml:"platform_type"`

	FlavorId *FlavorId `json:"flavor_id,omitempty" xml:"flavor_id"`

	// 组件部署件。key为组件component_name，对于Docker多容器场景，key为容器名称。
	Artifacts map[string]interface{} `json:"artifacts,omitempty" xml:"artifacts"`

	// 应用组件版本号。
	Version *string `json:"version,omitempty" xml:"version"`

	// 应用组件配置，如环境变量。
	Configuration *interface{} `json:"configuration,omitempty" xml:"configuration"`

	// 创建人。
	Creator *string `json:"creator,omitempty" xml:"creator"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	// 访问方式列表。
	ExternalAccesses *[]ExternalAccesses `json:"external_accesses,omitempty" xml:"external_accesses"`

	// 部署资源列表。
	ReferResources *[]ReferResources `json:"refer_resources,omitempty" xml:"refer_resources"`

	StatusDetail   *InstanceStatusView `json:"status_detail,omitempty" xml:"status_detail"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowInstanceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDetailResponse", string(data)}, " ")
}
