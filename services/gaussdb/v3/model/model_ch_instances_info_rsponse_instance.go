package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChInstancesInfoRsponseInstance 实例信息。
type ChInstancesInfoRsponseInstance struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`

	// 租户在某一Region下的project ID。
	ProjectId string `json:"project_id"`

	// 公网访问IP。
	PublicIp *string `json:"public_ip,omitempty"`

	// 是否可公网访问。
	CanEnablePublicAccess bool `json:"can_enable_public_access"`

	// 备份节点ID。
	CurrentBackupNodeId *string `json:"current_backup_node_id,omitempty"`

	// 部署模式。 取值范围： - Single：单机 - Ha：主备
	ClusterMode string `json:"cluster_mode"`

	// 实例状态。
	Status string `json:"status"`

	// 是否冻结。 取值范围： - 0：不冻结 - 1：冻结
	IsFrozen int32 `json:"is_frozen"`

	// 冻结时间。
	FrozenTime *string `json:"frozen_time,omitempty"`

	// 默认用户。
	DbUser *string `json:"db_user,omitempty"`

	// 备份周期。
	BakPeriod *string `json:"bak_period,omitempty"`

	// 备份保存天数。
	BakKeepDay *int32 `json:"bak_keep_day,omitempty"`

	// 备份预计开始时间。
	BakExpectedStartTime *string `json:"bak_expected_start_time,omitempty"`

	// 数据库版本ID。
	DatastoreVersionId string `json:"datastore_version_id"`

	// 数据库版本。
	DatastoreVersion string `json:"datastore_version"`

	// 数据库引擎。
	DatastoreType string `json:"datastore_type"`

	// 实例创建时间。
	CreateAt int64 `json:"create_at"`

	// 实例更新时间。
	UpdateAt int64 `json:"update_at"`

	// 实例删除时间。
	DeleteAt *int64 `json:"delete_at,omitempty"`

	// 数据库端口号。取值范围0~65535。
	DbPort string `json:"db_port"`

	ParamGroup *ChInstancesInfoRsponseInstanceParamGroup `json:"param_group,omitempty"`

	// 实例动作。
	Actions *[]ChQueryActionInfo `json:"actions,omitempty"`

	// 实例创建失败错误码。
	CreateFailErrorCode *string `json:"create_fail_error_code,omitempty"`

	// 实例分组。
	Groups []ChInstancesInfoRsponseInstanceGroups `json:"groups"`

	OpsWindow *ChInstancesInfoRsponseInstanceOpsWindow `json:"ops_window"`

	TagsInfo *CreateChInstanceInfoTagsInfo `json:"tags_info,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 备份使用空间。
	BackupUsedSpace *string `json:"backup_used_space,omitempty"`

	// 可用区模式。 取值范围： - single：单可用区 - double：多可用区-
	AzMode string `json:"az_mode"`

	// 主可用区。
	MasterAzCode *string `json:"master_az_code,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	PortInfo *ChInstancesInfoRsponseInstancePortInfo `json:"port_info"`

	// 磁盘规格码。
	VolumeCode string `json:"volume_code"`

	// 是否支持副本。
	SupportDataReplication bool `json:"support_data_replication"`

	// 是否有数据库新版本。
	NewVersionAvailable bool `json:"new_version_available"`

	// SSL开关。
	SslOption bool `json:"ssl_option"`

	// 专属资源池ID。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	// 支付模式。 取值范围： - 0：按需计费 - 1：包周期-
	PayModel *string `json:"pay_model,omitempty"`
}

func (o ChInstancesInfoRsponseInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstancesInfoRsponseInstance struct{}"
	}

	return strings.Join([]string{"ChInstancesInfoRsponseInstance", string(data)}, " ")
}
