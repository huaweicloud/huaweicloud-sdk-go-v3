package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterDetailInstance struct {
	Flavor *ClusterDetailInstanceFlavor `json:"flavor"`

	Volume *ClusterDetailInstanceVolume `json:"volume"`
	// 节点状态： - 100：创建中。 - 200：正常。 - 300：失败。 - 303：创建失败。 - 400：已删除。 - 800：冻结。

	Status string `json:"status"`
	// 节点操作状态列表： - REBOOTING：重启中。 - RESTORING：恢复中。 - REBOOT_FAILURE：重启失败。

	Actions *[]string `json:"actions,omitempty"`
	// 节点类型，只支持一种类型“cdm”。

	Type string `json:"type"`
	// 节点的虚拟机ID。

	Id string `json:"id"`
	// 节点的虚拟机名称。

	Name string `json:"name"`
	// 节点是否冻结：0：否。1：是。

	IsFrozen string `json:"isFrozen"`
	// 节点配置状态： - In-Sync：配置已同步。 - Applying：配置中。 - Sync-Failure：配置失败。

	ConfigStatus *string `json:"config_status,omitempty"`
	// 实例角色

	Role *string `json:"role,omitempty"`
	// 分组

	Group *string `json:"group,omitempty"`
	// 链接信息

	Links *[]ClusterLinks `json:"links,omitempty"`
	// 组件分组id

	ParamsGroupId *string `json:"paramsGroupId,omitempty"`
	// 公网ip

	PublicIp *string `json:"publicIp,omitempty"`
	// 管理ip

	ManageIp *string `json:"manageIp,omitempty"`
	// 流量ip

	TrafficIp *string `json:"trafficIp,omitempty"`
	// 分片id

	ShardId *string `json:"shard_id,omitempty"`
	// 管理修复ip

	ManageFixIp *string `json:"manage_fix_ip,omitempty"`
	// 私有ip

	PrivateIp *string `json:"private_ip,omitempty"`
	// 内部ip

	InternalIp *string `json:"internal_ip,omitempty"`
	// 资源信息

	Resource *[]Resource `json:"resource,omitempty"`
}

func (o ClusterDetailInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetailInstance struct{}"
	}

	return strings.Join([]string{"ClusterDetailInstance", string(data)}, " ")
}
