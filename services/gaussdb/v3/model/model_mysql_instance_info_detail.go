package model

import (
	"encoding/json"

	"strings"
)

type MysqlInstanceInfoDetail struct {
	// 实例ID。

	Id string `json:"id"`
	// 创建的实例名称。

	Name string `json:"name"`
	// 租户在某一region下的project ID。

	ProjectId string `json:"project_id"`
	// 实例状态。

	Status *string `json:"status,omitempty"`
	// 数据库端口号。

	Port *string `json:"port,omitempty"`
	// 实例类型，取值为“Cluster”。

	Type *string `json:"type,omitempty"`
	// 节点个数。

	NodeCount *int32 `json:"node_count,omitempty"`

	Datastore *MysqlDatastore `json:"datastore,omitempty"`
	// 备份空间使用大小，单位为GB。

	BackupUsedSpace float32 `json:"backup_used_space,omitempty"`
	// 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。说明：创建时返回值为空，数据库实例创建成功后该值不为空。

	Created *string `json:"created,omitempty"`
	// 更新时间，格式与\"created\"字段对应格式完全相同。说明：创建时返回值为空，数据库实例创建成功后该值不为空。

	Updated *string `json:"updated,omitempty"`
	// 实例的写内网IP。

	PrivateWriteIps *[]string `json:"private_write_ips,omitempty"`
	// 实例的公网IP。

	PublicIps *string `json:"public_ips,omitempty"`
	// 默认用户名。

	DbUserName *string `json:"db_user_name,omitempty"`
	// 虚拟私有云ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 子网的网络ID信息。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 安全组ID。

	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// 实例创建的模板ID，或者应用到实例的最新参数组模板ID。

	ConfigurationId *string `json:"configuration_id,omitempty"`

	Volume *MysqlVolumeInfo `json:"volume,omitempty"`

	BackupStrategy *MysqlBackupStrategy `json:"backup_strategy,omitempty"`

	Nodes *[]MysqlInstanceNodeInfo `json:"nodes,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 时区。

	TimeZone *string `json:"time_zone,omitempty"`
	// 可用区模式，单可用区single或多可用区multi。

	AzMode *string `json:"az_mode,omitempty"`
	// 主可用区。

	MasterAzCode *string `json:"master_az_code,omitempty"`
	// 可维护时间窗，为UTC时间。

	MaintenanceWindow *string `json:"maintenance_window,omitempty"`
	// 实例标签。

	Tags *[]MysqlTags `json:"tags,omitempty"`
}

func (o MysqlInstanceInfoDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlInstanceInfoDetail struct{}"
	}

	return strings.Join([]string{"MysqlInstanceInfoDetail", string(data)}, " ")
}
