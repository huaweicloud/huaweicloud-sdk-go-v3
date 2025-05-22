package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfo 集群列表对象。
type ClusterInfo struct {

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 集群名称。 **取值范围**： 同一个账号ID下唯一。
	Name string `json:"name"`

	// **参数解释**： 集群状态,字符串枚举。 **取值范围**： - CREATING：创建中 - ACTIVE：可用 - FAILED：不可用 - CREATE_FAILED：创建失败 - DELETING：删除中 - DELETE_FAILED：删除失败 - FROZEN：普通冻结 - POLICE_FROZEN：公安冻结
	Status string `json:"status"`

	// **参数解释**： 数据仓库集群版本。 **取值范围**： 小数点分割的3~4段字符串，如9.1.0.200，每一段数字越大版本越新。
	Version string `json:"version"`

	// **参数解释**： 集群上次修改时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ **取值范围**： 大于等于集群创建时间的ISO8601格式时间。
	Updated string `json:"updated"`

	// **参数解释**： 集群创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ **取值范围**： ISO8601格式的时间。
	Created string `json:"created"`

	// **参数解释**： 集群服务端口，创建集群时如未指定则默认8000。 **取值范围**： 8000~30000
	Port int32 `json:"port"`

	// **参数解释**： 集群的内网连接信息。 **取值范围**： 不涉及。
	Endpoints []Endpoints `json:"endpoints"`

	// **参数解释**： 集群实例。 **取值范围**： 列表大小与集群节点数量字段number_of_node相同。
	Nodes []Nodes `json:"nodes"`

	// **参数解释**： 集群标签。 **取值范围**： 默认null。
	Tags []Tags `json:"tags"`

	// **参数解释**： 管理员用户名。 **取值范围**： 默认dbadmin。
	UserName string `json:"user_name"`

	// **参数解释**： 节点数量。创建集群时指定。 **取值范围**： 不涉及。
	NumberOfNode int32 `json:"number_of_node"`

	// **参数解释**： 事件数。仅记录用户操作且对集群产生影响的事件，部分按钮开闭类操作不记入集群事件数。 **取值范围**： 不涉及。
	RecentEvent int32 `json:"recent_event"`

	// **参数解释**： 可用区。 **取值范围**： 不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释**： 企业项目ID，对集群指定企业项目。如果未指定，则使用默认企业项目“default”的ID，即0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释**： 集群规格ID。 **取值范围**： 不涉及。
	NodeType string `json:"node_type"`

	// **参数解释**： 虚拟私有云ID。 **取值范围**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 子网ID。 **取值范围**： 不涉及。
	SubnetId string `json:"subnet_id"`

	PublicIp *PublicIp `json:"public_ip"`

	// **参数解释**： 公网IP信息，如果未指定，则默认不支持公网连接。 **取值范围**： 不涉及。
	PublicEndpoints []PublicEndpoints `json:"public_endpoints"`

	// **参数解释**： 任务信息，由key、value组成。key值为正在进行的任务，value值为正在进行任务的进度。 **取值范围**： key值的有效值包括但不限于以下这些： - CREATING：创建中 - RESTORING：恢复中 - SNAPSHOTTING：快照中 - GROWING：扩容中 - REBOOTING：重启中 - SETTING_CONFIGURATION：安全设置配置中 - CONFIGURING_EXT_DATASOURCE：MRS连接配置中 - ADD_CN_ING：增加CN中 - DEL_CN_ING：删除CN中 - REDISTRIBUTING：重分布中 - ELB_BINDING：弹性负载均衡绑定中 - ELB_UNBINDING：弹性负载均衡解绑中 - ELB_SWITCHING：弹性负载均衡切换中 - NETWORK_CONFIGURING：网络配置中 - DISK_EXPANDING：磁盘扩容中 - ACTIVE_STANDY_SWITCHOVER：主备恢复中 - CLUSTER_SHRINKING：缩容中 - SHRINK_CHECKING：缩容检测中 - FLAVOR_RESIZING：规格变更中 - MANAGE_IP_BINDING：登录开通中 - FINE_GRAINED_RESTORING：细粒度恢复中 - DR_RECOVERING：容灾恢复中 - REPAIRING：修复中
	ActionProgress map[string]string `json:"action_progress"`

	// **参数解释**： “可用”集群状态的子状态。 **取值范围**： 有效值包括： - NORMAL：正常 - READONLY：只读 - REDISTRIBUTING：重分布中 - REDISTRIBUTION-FAILURE：重分布失败 - UNBALANCED：非均衡 - UNBALANCED | READONLY：非均衡，只读 - DEGRADED：节点故障 - DEGRADED | READONLY：节点故障，只读 - DEGRADED | UNBALANCED：节点故障，非均衡 - UNBALANCED | REDISTRIBUTING：非均衡，重分布中 - UNBALANCED | REDISTRIBUTION-FAILURE：非均衡，重分布失败 - READONLY | REDISTRIBUTION-FAILURE：只读，重分布失败 - UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：非均衡，只读，重分布失败 - DEGRADED | REDISTRIBUTION-FAILURE：节点故障，重分布失败 - DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE：节点故障，非均衡，只读，重分布失败 - DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：节点故障，非均衡，只读，重分布失败 - DEGRADED | UNBALANCED | READONLY：节点故障，非均衡，只读
	SubStatus string `json:"sub_status"`

	// **参数解释**： 集群管理任务，表示当前正在进行的任务或已执行的任务的结果。 **取值范围**： 有效值包括但不限于以下值： - UNFREEZING：解冻中 - FREEZING：冻结中 - RESTORING：恢复中 - SNAPSHOTTING：快照中 - GROWING：扩容中 - REBOOTING：重启中 - SETTING_CONFIGURATION：安全设置配置中 - CONFIGURING_EXT_DATASOURCE：MRS连接配置中 - DELETING_EXT_DATASOURCE：删除MRS连接 - REBOOT_FAILURE：重启失败 - RESIZE_FAILURE：扩容失败 - ADD_CN_ING：增加CN中 - DEL_CN_ING：删除CN中 - CREATING_NODE：添加节点 - CREATE_NODE_FAILED：添加节点失败 - DELETING_NODE：删除节点 - DELETE_NODE_FAILED：删除节点失败 - REDISTRIBUTING：重分布中 - REDISTRIBUTE_FAILURE：重分布失败 - WAITING_REDISTRIBUTION：待重分布 - REDISTRIBUTION_PAUSED：重分布暂停 - ELB_BINDING：弹性负载均衡绑定中 - ELB_BIND_FAILED：弹性负载均衡绑定失败 - ELB_UNBINDING：弹性负载均衡解绑中 - ELB_UNBIND_FAILED：弹性负载均衡解绑失败 - ELB_SWITCHING：弹性负载均衡切换中 - ELB_SWITCHING_FAILED：弹性负载均衡切换失败 - NETWORK_CONFIGURING：网络配置中 - NETWORK_CONFIG_FAILED：网络配置失败 - DISK_EXPAND_FAILED：磁盘扩容失败 - DISK_EXPANDING：磁盘扩容中 - ACTIVE_STANDY_SWITCHOVER：主备恢复中 - ACTIVE_STANDY_SWITCHOVER_FAILURE：主备恢复失败 - CLUSTER_SHRINK_FAILED：缩容失败 - CLUSTER_SHRINKING：缩容中 - SHRINK_CHECK_FAILED：缩容检测失败 - SHRINK_CHECKING：缩容检测中 - FLAVOR_RESIZING_FAILED：规格变更失败 - FLAVOR_RESIZING：规格变更中 - MANAGE_IP_BIND_FAILED：登录开通失败 - MANAGE_IP_BINDING：登录开通中 - ORDER_PENDING：订单待支付 - FINE_GRAINED_RESTORING：细粒度恢复中 - DR_RECOVERING：容灾恢复中
	TaskStatus string `json:"task_status"`

	// **参数解释**： 安全组ID。 **取值范围**： 不涉及。
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
