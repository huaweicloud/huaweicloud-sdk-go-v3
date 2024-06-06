package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleInstanceV3 struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例类型。
	HaMode *string `json:"ha_mode,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 计费模式。
	PayModel *string `json:"pay_model,omitempty"`

	// 创建时间。
	CreateAt *int32 `json:"create_at,omitempty"`

	// 删除时间。
	DeletedAt *int32 `json:"deleted_at,omitempty"`

	// 磁盘类型。
	VolumeType *string `json:"volume_type,omitempty"`

	// 磁盘大小。
	VolumeSize *string `json:"volume_size,omitempty"`

	// 数据面VIP。
	DataVip *string `json:"data_vip,omitempty"`

	// 数据面IPV6。
	DataVipIpv6 *string `json:"data_vip_ipv6,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 备份级别。
	BackupLevel *string `json:"backup_level,omitempty"`

	// 备份ID。
	RecycleBackupId *string `json:"recycle_backup_id,omitempty"`

	// 回收状态。
	RecycleStatus *string `json:"recycle_status,omitempty"`
}

func (o RecycleInstanceV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstanceV3 struct{}"
	}

	return strings.Join([]string{"RecycleInstanceV3", string(data)}, " ")
}
