package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlInstanceInfoDetailUnifyStatus struct {

	// 实例ID，严格匹配UUID规则。
	Id string `json:"id"`

	// 创建的实例名称。
	Name string `json:"name"`

	// 租户在某一Region下的project ID。
	ProjectId string `json:"project_id"`

	// 实例状态。  取值： - 值为“creating”，表示实例正在创建。 - 值为“normal”，表示实例正常。 - 值为“abnormal”，表示实例异常。 - 值为“createfail”，表示实例创建失败。
	Status *string `json:"status,omitempty"`

	// 数据库端口号。
	Port *string `json:"port,omitempty"`

	// 实例备注
	Alias *string `json:"alias,omitempty"`

	// 实例类型，取Cluster”。
	Type *string `json:"type,omitempty"`

	ChargeInfo *MysqlInstanceChargeInfo `json:"charge_info,omitempty"`

	// 节点个数。
	NodeCount *int32 `json:"node_count,omitempty"`

	Datastore *MysqlDatastoreWithKernelVersion `json:"datastore,omitempty"`

	// 备份空间使用大小，单位为GB。
	BackupUsedSpace *float64 `json:"backup_used_space,omitempty"`

	// 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式与\"created\"字段对应格式完全相同。
	Updated *string `json:"updated,omitempty"`

	// 实例的写内网IP地址。
	PrivateWriteIps *[]string `json:"private_write_ips,omitempty"`

	// 实例内网域名列表。实例创建成功后，需要手动申请内网域名，否则查询内网域名为空。
	PrivateDnsNames *[]string `json:"private_dns_names,omitempty"`

	// 实例的公网IP地址。
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

	BackupStrategy *MysqlBackupStrategy `json:"backup_strategy,omitempty"`

	// 节点信息。
	Nodes *[]MysqlInstanceNodeInfo `json:"nodes,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 可用区模式。  取值范围： - single：单可用区。 - multi：多可用区。
	AzMode *string `json:"az_mode,omitempty"`

	// 主可用区。
	MasterAzCode *string `json:"master_az_code,omitempty"`

	// 可维护时间窗，为UTC时间。
	MaintenanceWindow *string `json:"maintenance_window,omitempty"`

	// 实例标签。
	Tags *[]MysqlTags `json:"tags,omitempty"`

	// 专属资源池ID，只有数据库实例属于专属资源池才会返回该参数。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	// 代理信息。
	Proxies *[]MysqlProxyInfo `json:"proxies,omitempty"`

	TdeInfo *MysqlTdeInfo `json:"tde_info,omitempty"`
}

func (o MysqlInstanceInfoDetailUnifyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlInstanceInfoDetailUnifyStatus struct{}"
	}

	return strings.Join([]string{"MysqlInstanceInfoDetailUnifyStatus", string(data)}, " ")
}
