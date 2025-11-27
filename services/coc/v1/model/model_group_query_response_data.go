package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupQueryResponseData struct {

	// CloudCMDB分配的uuid。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 厂商（默认RMS，可填RMS/ALI/OTHER）。
	Vendor *string `json:"vendor,omitempty"`

	// code。
	Code *string `json:"code,omitempty"`

	// 租户id。
	DomainId *string `json:"domain_id,omitempty"`

	// region id。
	RegionId *string `json:"region_id,omitempty"`

	// component id。
	ComponentId *string `json:"component_id,omitempty"`

	// application id。
	ApplicationId *string `json:"application_id,omitempty"`

	// 企业项目id。
	EpId string `json:"ep_id"`

	// 资源关联模式，MANUAL表示手动关联，AUTO表示智能关联。
	SyncMode *string `json:"sync_mode,omitempty"`

	// 关联标签。
	RuleTags *string `json:"rule_tags,omitempty"`

	// 分组配置信息。
	RelationConfigurations *[]GroupRelationConfiguration `json:"relation_configurations,omitempty"`
}

func (o GroupQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupQueryResponseData struct{}"
	}

	return strings.Join([]string{"GroupQueryResponseData", string(data)}, " ")
}
