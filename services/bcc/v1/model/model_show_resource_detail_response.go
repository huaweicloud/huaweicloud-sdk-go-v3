package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceDetailResponse Response Object
type ShowResourceDetailResponse struct {

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务名称
	Provider *string `json:"provider,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// Openstack中的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// Openstack中的项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty"`

	// 资源详细属性
	ResourceDetail map[string]interface{} `json:"resource_detail,omitempty"`

	// 资源状态
	State *string `json:"state,omitempty"`

	ResourceLevel *ResourceLevelBase `json:"resource_level,omitempty"`

	// 等级是否手动指定
	IsLevelManual *bool `json:"is_level_manual,omitempty"`

	// 快照备份是否开启
	SnapshotBackupFlag *bool `json:"snapshot_backup_flag,omitempty"`

	// 本地备份是否开启
	LocalBackupFlag *bool `json:"local_backup_flag,omitempty"`

	// 异地备份是否开启
	RemoteBackupFlag *bool `json:"remote_backup_flag,omitempty"`

	ComplianceRule *ComplianceRule `json:"compliance_rule,omitempty"`

	// 资源保护状态
	InventoryResult *string `json:"inventory_result,omitempty"`

	// 本地备份存储库
	LocalVaults *[]string `json:"local_vaults,omitempty"`

	// 异地备份存储库
	RemoteVaults *[]string `json:"remote_vaults,omitempty"`

	// 资产合规盘点结果
	ComplianceResult *string `json:"compliance_result,omitempty"`

	// 资产合规盘点明细，合规校验针对local_backup_enabled,local_backup_frequency, local_worm_enabled, local_backup_retention 等每一项都给出规则要求值，资源实际值，是否满足匹配。结合多项给出最终的匹配结果。
	ComplianceDetail *string `json:"compliance_detail,omitempty"`

	// 标签
	Tags *[]TagItem `json:"tags,omitempty"`

	// 账户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 资源创建时间
	ResourceCreated *int64 `json:"resource_created,omitempty"`

	// 资源更新时间
	ResourceUpdated *int64 `json:"resource_updated,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowResourceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceDetailResponse", string(data)}, " ")
}
