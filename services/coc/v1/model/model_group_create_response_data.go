package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupCreateResponseData struct {

	// CloudCMDB分配的uuid。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 厂商（默认RMS，可填RMS/ALI/OTHER）。
	Vendor *string `json:"vendor,omitempty"`

	// 分组code。
	Code *string `json:"code,omitempty"`

	// 租户id。
	DomainId *string `json:"domain_id,omitempty"`

	// region id。
	RegionId *string `json:"region_id,omitempty"`

	// 组件id。
	ComponentId *string `json:"component_id,omitempty"`

	// 应用id。
	ApplicationId *string `json:"application_id,omitempty"`

	// 分组路径。
	Path *string `json:"path,omitempty"`

	// 资源关联模式，MANUAL表示手动关联，AUTO表示智能关联。
	SyncMode *string `json:"sync_mode,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 智能关联规则。
	SyncRules *[]GroupUpdateRequestSyncRules `json:"sync_rules,omitempty"`

	// 跨帐号资源所属的domainId。
	RelatedDomainId *string `json:"related_domain_id,omitempty"`

	// 分组配置信息。
	RelationConfigurations *[]GroupRelationConfiguration `json:"relation_configurations,omitempty"`
}

func (o GroupCreateResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupCreateResponseData struct{}"
	}

	return strings.Join([]string{"GroupCreateResponseData", string(data)}, " ")
}
