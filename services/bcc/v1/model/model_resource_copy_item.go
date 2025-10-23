package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceCopyItem 单条资源副本记录
type ResourceCopyItem struct {

	// 副本ID
	Id *string `json:"id,omitempty"`

	// 源服务的备份ID
	CopyId *string `json:"copy_id,omitempty"`

	// 副本名称
	CopyName *string `json:"copy_name,omitempty"`

	// 备份类型
	CopyType *string `json:"copy_type,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 可用区ID
	AzId *string `json:"az_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty"`

	// 副本当前状态
	Status *string `json:"status,omitempty"`

	// 备份开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 备份结束时间
	EndTime *string `json:"end_time,omitempty"`

	// CBR存储库ID
	VaultId *string `json:"vault_id,omitempty"`

	// CBR存储库名称
	VaultName *string `json:"vault_name,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResourceCopyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCopyItem struct{}"
	}

	return strings.Join([]string{"ResourceCopyItem", string(data)}, " ")
}
