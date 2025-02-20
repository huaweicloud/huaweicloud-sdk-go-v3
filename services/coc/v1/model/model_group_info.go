package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupInfo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Code *string `json:"code,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	RegionId *string `json:"region_id,omitempty"`

	ApplicationId *string `json:"application_id,omitempty"`

	ComponentId *string `json:"component_id,omitempty"`

	SyncMode *string `json:"sync_mode,omitempty"`

	Vendor *string `json:"vendor,omitempty"`

	SyncRules *string `json:"sync_rules,omitempty"`

	RelationConfigurations *[]string `json:"relation_configurations,omitempty"`
}

func (o GroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupInfo struct{}"
	}

	return strings.Join([]string{"GroupInfo", string(data)}, " ")
}
