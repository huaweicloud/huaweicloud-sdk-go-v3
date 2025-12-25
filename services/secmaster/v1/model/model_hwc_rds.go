package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcRds 云关系型数据库
type HwcRds struct {

	// 实例ID。
	Id string `json:"id"`

	// 创建的实例名称。
	Name string `json:"name"`

	// DBSS开启状态：OPEN | CLOSE
	ProtectedStatus string `json:"protected_status"`

	// 实例状态。 取值： 值为“BUILD”，表示实例正在创建。 值为“ACTIVE”，表示实例正常。 值为“FAILED”，表示实例异常。 值为“FROZEN”，表示实例冻结。 值为“MODIFYING”，表示实例正在扩容。 值为“REBOOTING”，表示实例正在重启。 值为“RESTORING”，表示实例正在恢复。 值为“MODIFYING INSTANCE TYPE”，表示实例正在转主备。 值为“SWITCHOVER”，表示实例正在主备切换。 值为“MIGRATING”，表示实例正在迁移。 值为“BACKING UP”，表示实例正在进行备份。 值为“MODIFYING DATABASE PORT”，表示实例正在修改数据库端口。 值为“STORAGE FULL”，表示实例磁盘空间满。
	Status string `json:"status"`

	// 实例的备注信息。
	Alias *string `json:"alias,omitempty"`

	// 实例内网IP地址列表。弹性云服务器创建成功后该值存在，其他情况下为空字符串。
	PrivateIps *[]string `json:"private_ips,omitempty"`

	// 实例内网域名列表。实例创建成功后，需要手动申请内网域名，否则查询内网域名为空。
	PrivateDnsNames *[]string `json:"private_dns_names,omitempty"`

	// 实例外网IP地址列表。
	PublicIps *[]string `json:"public_ips,omitempty"`

	// 数据库端口号。 RDS for MySQL数据库端口设置范围为1024～65535（其中12017和33071被RDS系统占用不可设置）。 RDS for PostgreSQL数据库端口修改范围为2100～9500。 RDS for SQL Server实例的端口设置范围为1433和2100~9500（其中5355和5985不可设置。对于2017 EE、2017 SE、2017 Web版，5050、5353和5986不可设置）。 当不传该参数时，默认端口如下：  RDS for MySQL默认3306。 RDS for PostgreSQL默认5432。 RDS for SQL Server默认1433。
	Port int32 `json:"port"`

	// 实例开启SSL标志。 取值为“true”：表示实例已开启SSL。 取值为“false”：表示实例未开启SSL。
	EnableSsl bool `json:"enable_ssl"`

	// 实例类型，取值为“Single”，“Ha”或“Replica”, \"Enterprise\"，分别对应于单机实例、主备实例、只读实例、分布式实例（企业版）。
	Type string `json:"type"`

	Ha *HwcRdsHa `json:"ha"`

	// 实例所在区域。
	Region string `json:"region"`

	Datastore *HwcRdsDatastore `json:"datastore"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。 说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Created string `json:"created"`

	// 更新时间，格式与“created”字段对应格式完全相同。 说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Updated *string `json:"updated,omitempty"`

	// 默认用户名。
	DbUserName *string `json:"db_user_name,omitempty"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 子网的网络ID信息。
	SubnetId string `json:"subnet_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 规格码。
	FlavorRef string `json:"flavor_ref"`

	// CPU大小。例如，1表示1U。
	Cpu string `json:"cpu"`

	// 内存大小（单位：GB）。
	Mem string `json:"mem"`

	Volume *HwcRdsVolume `json:"volume"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目标签ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 数据库切换策略。取值为“reliability”或“availability”，分别对应于可靠性优先和可用性优先。
	SwitchStrategy string `json:"switch_strategy"`

	// 用户设置的实例只读状态。仅支持RDS for MySQL引擎。 true，表示该实例被设置为只读状态。 false，表示该实例未被设置为只读状态。
	ReadOnlyByUser bool `json:"read_only_by_user"`

	BackupStrategy *HwcRdsBackupStrategy `json:"backup_strategy"`

	// 可维护时间窗，为UTC时间。
	MaintenanceWindow string `json:"maintenance_window"`

	// 主备实例信息
	Nodes []HwcRdsNode `json:"nodes"`

	// 所关联的数据库实例列表。
	RelatedInstance []HwcRdsRelatedInstance `json:"related_instance"`

	// 磁盘加密密钥ID。
	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	// 时区。
	TimeZone string `json:"time_zone"`

	// 备份空间使用量，单位GB。 该字段仅用于查询指定RDS for SQL Server单个实例信息时返回。
	BackupUsedSpace *float64 `json:"backup_used_space,omitempty"`

	// 磁盘空间使用量，单位GB。 该字段仅用于查询指定RDS for SQL Server单个实例信息时返回。
	StorageUsedSpace float64 `json:"storage_used_space"`

	// 是否已被DDM实例关联。
	AssociatedWithDdm bool `json:"associated_with_ddm"`

	// 实例磁盘的最大IOPS值。 当前该字段仅对于SQL Server引擎实例返回。
	MaxIops int64 `json:"max_iops"`

	// 实例的到期时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 仅包周期场景返回。
	ExpirationTime string `json:"expiration_time"`
}

func (o HwcRds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcRds struct{}"
	}

	return strings.Join([]string{"HwcRds", string(data)}, " ")
}
