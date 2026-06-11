package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesResult **参数解释：** 实例信息。 **取值范围：** 不涉及。
type ListInstancesResult struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 实例状态。 **取值范围：** - normal，表示实例正常。 - abnormal，表示实例异常。 - creating，表示实例创建中。 - frozen，表示实例被冻结。 - data_disk_full，表示实例磁盘已满。 - createfail，表示实例创建失败。 - enlargefail，表示实例扩容节点个数失败。
	Status string `json:"status"`

	// **参数解释：** 数据库端口。 **取值范围：** 不涉及。
	Port string `json:"port"`

	// **参数解释：** 实例所在区域。 **取值范围：** 不涉及。
	Region string `json:"region"`

	Datastore *ListInstancesDatastoreResult `json:"datastore"`

	// **参数解释：** 实例类型。 **取值范围：** 与请求参数相同。
	Mode string `json:"mode"`

	// **参数解释：** 产品类型。 **取值范围：** GeminiDB Redis云原生部署模式集群涉及此字段，取值：   -  Standard 标准型   -  Capacity 容量型
	ProductType *string `json:"product_type,omitempty"`

	// **参数解释：** 存储引擎。 **取值范围：** 取值为“rocksDB”。
	Engine string `json:"engine"`

	// **参数解释：** 实例创建时间。 **取值范围：** 不涉及。
	Created string `json:"created"`

	// **参数解释：** 实例操作最新变更的时间。 **取值范围：** 不涉及。
	Updated string `json:"updated"`

	// **参数解释：** 默认用户名。 **取值范围：** 取值为“rwuser”。
	DbUserName string `json:"db_user_name"`

	// **参数解释：** 虚拟私有云ID。 **取值范围：** 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释：** 子网ID。 **取值范围：** GeminiDB Cassandra 实例使用多个子网的场景，请参见表 ListInstancesNodeResult字段数据结构说明中的“subnet_id”。
	SubnetId string `json:"subnet_id"`

	// **参数解释：** 安全组ID。 **取值范围：** 不涉及。
	SecurityGroupId string `json:"security_group_id"`

	BackupStrategy *ListInstancesBackupStrategyResult `json:"backup_strategy"`

	// **参数解释：** 计费方式。 **取值范围：** - 取值为“0”，表示按需计费。 - 取值为“1”，表示包年/包月计费。
	PayMode string `json:"pay_mode"`

	// **参数解释：** 系统可维护时间窗。 **取值范围：** 不涉及。
	MaintenanceWindow string `json:"maintenance_window"`

	// **参数解释：** 组信息。 **取值范围：** 不涉及。
	Groups []ListInstancesGroupResult `json:"groups"`

	// **参数解释：** 企业项目ID。 **取值范围：** 取值为“0”，表示为default企业项目。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释：** 专属资源ID。只有数据库实例属于专属资源池才会返回该参数。 **取值范围：** 不涉及。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	// **参数解释：** 时区。 **取值范围：** 不涉及。
	TimeZone string `json:"time_zone"`

	// **参数解释：** 实例正在执行的动作。 **取值范围：** 不涉及。
	Actions []string `json:"actions"`

	// **参数解释：** 磁盘加密时的密钥ID。 **取值范围：** 不涉及。
	DiskEncryptionId string `json:"disk_encryption_id"`

	// **参数解释：** 负载均衡ip。 **取值范围：** 只有存在负载均衡ip，才会返回该参数。
	LbIpAddress *string `json:"lb_ip_address,omitempty"`

	// **参数解释：** 负载均衡端口。 **取值范围：** 只有存在负载均衡ip，才会返回该参数。
	LbPort *string `json:"lb_port,omitempty"`

	// **参数解释：** 实例可用区。 **取值范围：** 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释：** 容灾实例ID。 **取值范围：** 不涉及。
	DrInstanceId *string `json:"dr_instance_id,omitempty"`

	DualActiveInfo *DualActiveInfo `json:"dual_active_info,omitempty"`

	CcmCertInfo *CertInfoOption `json:"ccm_cert_info,omitempty"`

	// **参数解释：** SSL安全连接启用情况。 **取值范围：** - 取值为“0”表示未启用。 - 取值为“1”表示已启用。
	Ssl string `json:"ssl"`

	BackupSpaceUsage *BackupSpaceUsage `json:"backup_space_usage,omitempty"`
}

func (o ListInstancesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResult struct{}"
	}

	return strings.Join([]string{"ListInstancesResult", string(data)}, " ")
}
