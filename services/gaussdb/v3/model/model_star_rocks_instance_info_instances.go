package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StarRocksInstanceInfoInstances struct {

	// 实例ID，严格匹配UUID规则。
	Id *string `json:"id,omitempty"`

	// 创建的实例名称。
	Name *string `json:"name,omitempty"`

	// 租户在某一Region下的project ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 公网访问IP。
	PublicIp *string `json:"public_ip,omitempty"`

	// StarRocks FE节点类型数据IP（多个IP使用逗号分隔）
	DataVip *string `json:"data_vip,omitempty"`

	// 是否可公网访问。
	CanEnablePublicAccess *bool `json:"can_enable_public_access,omitempty"`

	// 备份节点ID。
	CurrentBackupNodeId *string `json:"current_backup_node_id,omitempty"`

	// 部署模式。
	ClusterMode *string `json:"cluster_mode,omitempty"`

	// 实例状态。
	Status *string `json:"status,omitempty"`

	// 是否冻结。
	IsFrozen *int32 `json:"is_frozen,omitempty"`

	// 冻结时间。
	FrozenTime *int64 `json:"frozen_time,omitempty"`

	// 默认用户名。
	DbUser *string `json:"db_user,omitempty"`

	// 备份周期。
	BakPeriod *string `json:"bak_period,omitempty"`

	// 备份保存天数。
	BakKeepDay *int32 `json:"bak_keep_day,omitempty"`

	// 备份预计开始时间。
	BakExpectedStartTime *int64 `json:"bak_expected_start_time,omitempty"`

	// 数据库版本ID。
	DataStoreVersionId *string `json:"data_store_version_id,omitempty"`

	// 数据库版本。
	DataStoreVersion *string `json:"data_store_version,omitempty"`

	// 数据库引擎。
	DataStoreType *string `json:"data_store_type,omitempty"`

	// 实例创建时间。
	CreateAt *int64 `json:"create_at,omitempty"`

	// 实例更新时间。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// 实例删除时间。
	DeleteAt *int64 `json:"delete_at,omitempty"`

	// 数据库端口号。
	DbPort *string `json:"db_port,omitempty"`

	// 参数组。
	ParamGroup *string `json:"param_group,omitempty"`

	// 实例动作。
	Actions *[]QueryAction `json:"actions,omitempty"`

	// 实例创建失败错误码。
	CreateFailErrorCode *string `json:"create_fail_error_code,omitempty"`

	// 实例分组。
	Groups *[]StarRocksInstanceInfoGroups `json:"groups,omitempty"`

	OpsWindow *StarRocksInstanceInfoOpsWindow `json:"ops_window,omitempty"`

	TagsInfo *StarRocksInstanceInfoTagsInfo `json:"tags_info,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 备份使用空间。
	BackupUsedSpace *string `json:"backup_used_space,omitempty"`

	// 可用区模式。  取值范围：  - single：单可用区 - multi：多可用区
	AzMode *string `json:"az_mode,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	PortInfo *StarRocksInstanceInfoPortInfo `json:"port_info,omitempty"`

	// FE节点磁盘类型。
	FeNodeVolumeCode *string `json:"fe_node_volume_code,omitempty"`

	// BE节点磁盘类型。
	BeNodeVolumeCode *string `json:"be_node_volume_code,omitempty"`

	// FE节点磁盘大小。
	FeNodeVolumeSize *string `json:"fe_node_volume_size,omitempty"`

	// BE节点磁盘大小。
	BeNodeVolumeSize *string `json:"be_node_volume_size,omitempty"`

	// 是否支持副本。
	SupportDataReplication *bool `json:"support_data_replication,omitempty"`

	// 是否有数据库新版本。
	NewVersionAvailable *bool `json:"new_version_available,omitempty"`

	// SSL开关。
	SslOption *bool `json:"ssl_option,omitempty"`

	// 专属资源池ID，只有数据库实例属于专属资源池才会返回该参数。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	// 支付模式。
	PayModel *string `json:"pay_model,omitempty"`
}

func (o StarRocksInstanceInfoInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoInstances struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoInstances", string(data)}, " ")
}
