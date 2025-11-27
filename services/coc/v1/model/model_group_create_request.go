package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupCreateRequest struct {

	// 分组名称。
	Name string `json:"name"`

	// 组件id。
	ComponentId string `json:"component_id"`

	// 厂商（默认RMS，可填RMS/ALI/OTHER）。
	Vendor *string `json:"vendor,omitempty"`

	// region id。
	RegionId string `json:"region_id"`

	// 应用id。
	ApplicationId *string `json:"application_id,omitempty"`

	// 资源同步方式，MANUAL表示手动，AUTO表示智能关联。
	SyncMode string `json:"sync_mode"`

	// 智能关联规则。
	SyncRules *[]GroupUpdateRequestSyncRules `json:"sync_rules,omitempty"`

	// 分组配置信息。
	RelationConfigurations *[]GroupRelationConfiguration `json:"relation_configurations,omitempty"`
}

func (o GroupCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupCreateRequest struct{}"
	}

	return strings.Join([]string{"GroupCreateRequest", string(data)}, " ")
}
