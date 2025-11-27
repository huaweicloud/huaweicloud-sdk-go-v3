package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupUpdateRequest struct {

	// 分组名称。
	Name string `json:"name"`

	// 资源关联方式，MANUAL表示手动，AUTO表示智能。
	SyncMode string `json:"sync_mode"`

	// 智能关联规则。
	SyncRules *[]GroupUpdateRequestSyncRules `json:"sync_rules,omitempty"`

	// 分组配置信息。
	RelationConfigurations *[]GroupRelationConfiguration `json:"relation_configurations,omitempty"`
}

func (o GroupUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupUpdateRequest struct{}"
	}

	return strings.Join([]string{"GroupUpdateRequest", string(data)}, " ")
}
