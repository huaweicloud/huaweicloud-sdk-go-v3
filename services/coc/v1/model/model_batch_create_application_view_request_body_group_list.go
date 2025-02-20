package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodyGroupList struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// region id
	RegionId *string `json:"region_id,omitempty"`

	// 父节点code
	ParentName *string `json:"parent_name,omitempty"`

	// 同步模式
	SyncMode *string `json:"sync_mode,omitempty"`

	// 关联的资源id列表
	CmdbResourceIdList *[]string `json:"cmdb_resource_id_list,omitempty"`

	// 智能关联规则
	SyncRules *[]BatchCreateApplicationViewRequestBodySyncRules `json:"sync_rules,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodyGroupList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodyGroupList struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodyGroupList", string(data)}, " ")
}
