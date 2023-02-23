package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAggregateResourceConfigResponse struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 聚合器ID。
	AggregatorId *string `json:"aggregator_id,omitempty"`

	// 聚合器帐号。
	AggregatorDomainId *string `json:"aggregator_domain_id,omitempty"`

	// 聚合资源所属帐号的ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 企业项目ID。
	EpId *string `json:"ep_id,omitempty"`

	// 云服务名称。
	Provider *string `json:"provider,omitempty"`

	// 资源类型。
	Type *string `json:"type,omitempty"`

	// 资源名称。
	Name *string `json:"name,omitempty"`

	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`

	// Openstack中的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 资源创建时间。
	Created *string `json:"created,omitempty"`

	// 资源更新时间。
	Updated *string `json:"updated,omitempty"`

	// 资源标签。
	Tags map[string]string `json:"tags,omitempty"`

	// 资源详细属性。
	Properties     map[string]interface{} `json:"properties,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowAggregateResourceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregateResourceConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAggregateResourceConfigResponse", string(data)}, " ")
}
