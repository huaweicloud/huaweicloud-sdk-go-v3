package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfo 集群列表对象。
type ClusterInfo struct {

	// 集群ID。
	Id string `json:"id"`

	// 集群名称
	Name string `json:"name"`

	// 集群状态，有效值包括： - CREATING：创建中 - ACTIVE：可用 - FAILED：不可用 - CREATE_FAILED：创建失败 - DELETING：删除中 - DELETE_FAILED：删除失败 - DELETED：已删除 - FROZEN：普通冻结 - POLICE_FROZEN：公安冻结
	Status string `json:"status"`

	// 数据仓库版本
	Version string `json:"version"`

	// 集群上次修改时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ
	Updated string `json:"updated"`

	// 集群创建时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ。
	Created string `json:"created"`

	// 集群服务端口。
	Port int32 `json:"port"`

	// 集群的内网连接信息
	Endpoints []Endpoints `json:"endpoints"`

	// 集群实例
	Nodes []Nodes `json:"nodes"`

	// 集群标签
	Tags []Tags `json:"tags"`

	// 管理员用户名
	UserName string `json:"user_name"`

	// 节点数量
	NumberOfNode int32 `json:"number_of_node"`

	// 事件数
	RecentEvent int32 `json:"recent_event"`

	// 可用区
	AvailabilityZone string `json:"availability_zone"`

	// 企业项目ID。值为0表示默认企业项目“default”
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 节点类型
	NodeType string `json:"node_type"`

	// 虚拟私有云ID
	VpcId string `json:"vpc_id"`

	// 子网ID
	SubnetId string `json:"subnet_id"`

	PublicIp *PublicIp `json:"public_ip"`

	// 集群的公网连接信息，如果未指定，则默认不使用公网连接信息。
	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`

	// 任务信息，由key、value组成。key值为正在进行的任务，value值为正在进行任务的进度。key值的有效值包括： - CREATING：创建中 - RESTORING：恢复中 - SNAPSHOTTING：快照中 - GROWING：扩容中 - REBOOTING：重启中 - SETTING_CONFIGURATION：安全设置配置中 - CONFIGURING_EXT_DATASOURCE：MRS连接配置中 - ADD_CN_ING：增加CN中 - DEL_CN_ING：删除CN中 - REDISTRIBUTING：重分布中 - ELB_BINDING：弹性负载均衡绑定中 - ELB_UNBINDING：弹性负载均衡解绑中 - ELB_SWITCHING：弹性负载均衡切换中 - NETWORK_CONFIGURING：网络配置中 - DISK_EXPANDING：磁盘扩容中 - ACTIVE_STANDY_SWITCHOVER：主备恢复中 - CLUSTER_SHRINKING：缩容中 - SHRINK_CHECKING：缩容检测中 - FLAVOR_RESIZING：规格变更中 - MANAGE_IP_BINDING：登录开通中 - FINE_GRAINED_RESTORING：细粒度恢复中 - DR_RECOVERING：容灾恢复中 - REPAIRING：修复中
	ActionProgress map[string]string `json:"action_progress"`

	// “可用”集群状态的子状态，有效值包括：  - NORMAL：正常 - READONLY：只读 - REDISTRIBUTING：重分布中 - REDISTRIBUTION-FAILURE：重分布失败 - UNBALANCED：非均衡 - UNBALANCED | READONLY：非均衡，只读 - DEGRADED：节点故障 - DEGRADED | READONLY：节点故障，只读 - DEGRADED | UNBALANCED：节点故障，非均衡 - UNBALANCED | REDISTRIBUTING：非均衡，重分布中 - UNBALANCED | REDISTRIBUTION-FAILURE：非均衡，重分布失败 - READONLY | REDISTRIBUTION-FAILURE：只读，重分布失败 - UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：非均衡，只读，重分布失败 - DEGRADED | REDISTRIBUTION-FAILURE：节点故障，重分布失败 - DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE：节点故障，非均衡，只读，重分布失败 - DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：节点故障，非均衡，只读，重分布失败 - DEGRADED | UNBALANCED | READONLY：节点故障，非均衡，只读
	SubStatus string `json:"sub_status"`

	// 集群管理任务，有效值包括： - UNFREEZING：解冻中 - FREEZING：冻结中 - RESTORING：恢复中 - SNAPSHOTTING：快照中 - GROWING：扩容中 - REBOOTING：重启中 - SETTING_CONFIGURATION：安全设置配置中 - CONFIGURING_EXT_DATASOURCE：MRS连接配置中 - DELETING_EXT_DATASOURCE：删除MRS连接 - REBOOT_FAILURE：重启失败 - RESIZE_FAILURE：扩容失败 - ADD_CN_ING：增加CN中 - DEL_CN_ING：删除CN中 - CREATING_NODE：添加节点 - CREATE_NODE_FAILED：添加节点失败 - DELETING_NODE：删除节点 - DELETE_NODE_FAILED：删除节点失败 - REDISTRIBUTING：重分布中 - REDISTRIBUTE_FAILURE：重分布失败 - WAITING_REDISTRIBUTION：待重分布 - REDISTRIBUTION_PAUSED：重分布暂停 - ELB_BINDING：弹性负载均衡绑定中 - ELB_BIND_FAILED：弹性负载均衡绑定失败 - ELB_UNBINDING：弹性负载均衡解绑中 - ELB_UNBIND_FAILED：弹性负载均衡解绑失败 - ELB_SWITCHING：弹性负载均衡切换中 - ELB_SWITCHING_FAILED：弹性负载均衡切换失败 - NETWORK_CONFIGURING：网络配置中 - NETWORK_CONFIG_FAILED：网络配置失败 - DISK_EXPAND_FAILED：磁盘扩容失败 - DISK_EXPANDING：磁盘扩容中 - ACTIVE_STANDY_SWITCHOVER：主备恢复中 - ACTIVE_STANDY_SWITCHOVER_FAILURE：主备恢复失败 - CLUSTER_SHRINK_FAILED：缩容失败 - CLUSTER_SHRINKING：缩容中 - SHRINK_CHECK_FAILED：缩容检测失败 - SHRINK_CHECKING：缩容检测中 - FLAVOR_RESIZING_FAILED：规格变更失败 - FLAVOR_RESIZING：规格变更中 - MANAGE_IP_BIND_FAILED：登录开通失败 - MANAGE_IP_BINDING：登录开通中 - ORDER_PENDING：订单待支付 - FINE_GRAINED_RESTORING：细粒度恢复中 - DR_RECOVERING：容灾恢复中
	TaskStatus string `json:"task_status"`

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`

	FailedReasons *FailedReason `json:"failed_reasons,omitempty"`
}

func (o ClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfo", string(data)}, " ")
}
