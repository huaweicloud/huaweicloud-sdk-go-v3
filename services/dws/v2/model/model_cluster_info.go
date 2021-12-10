package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群列表对象
type ClusterInfo struct {
	// 任务信息，由key、value组成。key值为正在进行的任务，value值为正在进行任务的进度。key值的有效值包括：  GROWING：扩容中  RESTORING：恢复中  SNAPSHOTTING：快照中  REPAIRING : 修复中  CREATING : 创建中

	ActionProgress map[string]string `json:"action_progress"`

	FailedReasons *FailedReason `json:"failed_reasons,omitempty"`
	// 可用区

	AvailabilityZone string `json:"availability_zone"`
	// 集群的内网连接信息

	Endpoints []Endpoints `json:"endpoints"`
	// 集群管理任务，有效值包括：  RESTORING：恢复中  SNAPSHOTTING：快照中  GROWING：扩容中  REBOOTING：重启中  SETTING_CONFIGURATION：安全设置配置中  CONFIGURING_EXT_DATASOURCE：MRS连接配置中  DELETING_EXT_DATASOURCE：删除MRS连接  REBOOT_FAILURE：重启失败  RESIZE_FAILURE：扩容失败

	TaskStatus string `json:"task_status"`

	PublicIp *PublicIp `json:"public_ip"`
	// “可用”集群状态的子状态，有效值包括：  NORMAL：正常  READONLY：只读  REDISTRIBUTING：重分布中  REDISTRIBUTION-FAILURE：重分布失败  UNBALANCED：低性能  UNBALANCED | READONLY：低性能，只读  DEGRADED：节点故障  DEGRADED | READONLY：节点故障，只读  DEGRADED | UNBALANCED：节点故障，低性能  UNBALANCED | REDISTRIBUTING：低性能，重分布中  UNBALANCED | REDISTRIBUTION-FAILURE：低性能，重分布失败  READONLY | REDISTRIBUTION-FAILURE：只读，重分布失败  UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：低性能，只读，重分布失败  DEGRADED | REDISTRIBUTION-FAILURE：节点故障，重分布失败  DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE：节点故障，低性能，只读，重分布失败  DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：节点故障，低性能，只读，重分布失败  DEGRADED | UNBALANCED | READONLY：节点故障，低性能，只读

	SubStatus string `json:"sub_status"`
	// 节点数量

	NumberOfNode int32 `json:"number_of_node"`
	// 事件数

	RecentEvent int32 `json:"recent_event"`
	// 虚拟私有云ID

	VpcId string `json:"vpc_id"`
	// 集群创建时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ。

	Created string `json:"created"`
	// 管理员用户名

	UserName string `json:"user_name"`
	// 安全组ID

	SecurityGroupId string `json:"security_group_id"`
	// 数据仓库版本

	Version string `json:"version"`
	// 企业项目ID。值为0表示默认企业项目“default”

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 节点类型

	NodeType string `json:"node_type"`
	// 集群服务端口，取值范围8000~30000，默认值：8000

	Port int32 `json:"port"`
	// 集群名称

	Name string `json:"name"`
	// 子网ID

	SubnetId string `json:"subnet_id"`
	// 集群的公网连接信息，如果未指定，则默认不使用公网连接信息。

	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`
	// 集群ID

	Id string `json:"id"`
	// 集群上次修改时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ

	Updated string `json:"updated"`
	// 集群状态，有效值包括：  CREATING：创建中  AVAILABLE：可用  UNAVAILABLE：不可用  CREATION FAILED：创建失败

	Status string `json:"status"`
}

func (o ClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfo", string(data)}, " ")
}
