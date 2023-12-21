package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDetailStatusInfo 状态信息。
type InstanceDetailStatusInfo struct {

	// 云堡垒机实例状态。 - SHUTOFF：已关闭 - ACTIVE：运行中 - DELETING：删除中 - BUILD：创建中 - DELETED：已删除 - ERROR：故障 - HAWAIT：等待备机创建成功 - FROZEN：已冻结 - UPGRADING：升级中 - UNPAID：待支付 - RESIZE：规格变更中 - DILATATION：扩容中 - HA：配置HA中
	Status string `json:"status"`

	// 云堡垒机实例当前的任务状态。 - powering-on：开启 - powering-off：关闭 - rebooting：重启 - delete_wait：删除 - frozen：冻结 - NO_TASK：运行 - unfrozen：解冻 - alter：变更 - updating：升级中 - configuring-ha：配置HA - data-migrating：数据迁移中 - rollback：版本回滚中 - traffic-switchover：流量切换中
	TaskStatus string `json:"task_status"`

	// 云堡垒机实例在创建实例过程中的状态信息。 - Waiting for payment：等待支付 - creating-network：创建网络 - creating-server：创建服务 - tranfering-horizontal-network：网络打通 - adding-policy-route：添加路由策略 - configing-dns：配置DNS - starting-cbs-service：服务运行中 - setting-init-conf：初始化 - buying-EIP：购买弹性公网IP
	CreateInstanceStatus string `json:"create_instance_status"`

	// 云堡垒机实例状态。 - building：创建中 - deleting：删除中 - deleted：删除了 - unpaid：未支付 - upgrading：升级中 - resizing：扩容中 - abnormal：异常 - error：故障 - ok：正常
	InstanceStatus string `json:"instance_status"`

	// 云堡垒机实例信息描述。
	InstanceDescription string `json:"instance_description"`

	// 云堡垒机实例创建实例失败原因。
	FailReason string `json:"fail_reason"`
}

func (o InstanceDetailStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetailStatusInfo struct{}"
	}

	return strings.Join([]string{"InstanceDetailStatusInfo", string(data)}, " ")
}
