package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesResult 实例信息。
type ListInstancesResult struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`

	// 实例状态。 取值： - normal，表示实例正常。 - abnormal，表示实例异常。 - creating，表示实例创建中。 - frozen，表示实例被冻结。 - data_disk_full，表示实例磁盘已满。 - createfail，表示实例创建失败。 - enlargefail，表示实例扩容节点个数失败。
	Status string `json:"status"`

	// 数据库端口。
	Port string `json:"port"`

	// 实例类型。与请求参数相同。
	Mode string `json:"mode"`

	// 实例所在区域。
	Region string `json:"region"`

	Datastore *ListInstancesDatastoreResult `json:"datastore"`

	// 存储引擎。取值为“rocksDB”。
	Engine string `json:"engine"`

	// 实例创建时间。
	Created string `json:"created"`

	// 实例操作最新变更的时间。
	Updated string `json:"updated"`

	// 默认用户名。取值为“rwuser”。
	DbUserName string `json:"db_user_name"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	BackupStrategy *ListInstancesBackupStrategyResult `json:"backup_strategy"`

	// 计费方式。 - 取值为“0”，表示按需计费。 - 取值为“1”，表示包年/包月计费。
	PayMode string `json:"pay_mode"`

	// 系统可维护时间窗。
	MaintenanceWindow string `json:"maintenance_window"`

	// 组信息。
	Groups []ListInstancesGroupResult `json:"groups"`

	// 企业项目ID。取值为“0”，表示为default企业项目。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 专属资源ID，只有数据库实例属于专属资源池才会返回该参数。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	// 时区。
	TimeZone string `json:"time_zone"`

	// 实例正在执行的动作。
	Actions []string `json:"actions"`

	// 负载均衡ip，只有存在负载均衡ip，才会返回该参数。
	LbIpAddress *string `json:"lb_ip_address,omitempty"`

	// 负载均衡端口，只有存在负载均衡ip，才会返回该参数。
	LbPort *string `json:"lb_port,omitempty"`

	// 实例可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o ListInstancesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResult struct{}"
	}

	return strings.Join([]string{"ListInstancesResult", string(data)}, " ")
}
