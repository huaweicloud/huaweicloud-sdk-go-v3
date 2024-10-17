package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditInstanceListBean 审计实例列表bean
type AuditInstanceListBean struct {

	// 付费模式  - Period：包周期 - Demand：按需。
	ChargeModel string `json:"charge_model"`

	// 备注信息。
	Comment string `json:"comment"`

	// 配置的数据库总数。
	ConfigNum int32 `json:"config_num"`

	// 连接地址。
	ConnectIp string `json:"connect_ip"`

	// ipv6连接地址。
	ConnectIpv6 string `json:"connect_ipv6"`

	// CPU个数
	Cpu int32 `json:"cpu"`

	// 创建时间
	Created string `json:"created"`

	// 支持的数据库总数
	DatabaseLimit int32 `json:"database_limit"`

	// 实例结果状态 - 1:冻结可释放  - 2:冻结不可释放 - 3:冻结后不可续费
	Effect int32 `json:"effect"`

	// 过期时间
	Expired string `json:"expired"`

	// ID
	Id string `json:"id"`

	// 剩余天数
	KeepDays string `json:"keep_days"`

	// 实例别名
	Name string `json:"name"`

	// 如果有返回，则需要升级，如果没有，则为null。
	NewVersion string `json:"new_version"`

	// 绑定弹性IP的port ID
	PortId string `json:"port_id"`

	// 内存
	Ram int32 `json:"ram"`

	// 实例所在region
	Region string `json:"region"`

	// 到期天数
	RemainDays string `json:"remain_days"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 实例的规格
	ResourceSpecCode string `json:"resource_spec_code"`

	// 场景
	Scene string `json:"scene"`

	// 安全组
	SecurityGroupId string `json:"security_group_id"`

	// 实例规格
	Specification string `json:"specification"`

	// 实例状态  - SHUTOFF :已关闭  - ACTIVE: 运行中，允许任何操作   - DELETING: 删除中，不允许任何操作  - BUILD: 创建中，不允许任何操作  - DELETED: 已删除，不需要展示  - ERROR: 故障，只允许删除  - HAWAIT: 等待备机创建成功，不允许任何操作  - FROZEN: 已冻结，只允许续费、绑定/解绑  - UPGRADING: 升级中，不允许升级操作
	Status string `json:"status"`

	// 子网ID
	SubnetId string `json:"subnet_id"`

	// 任务状态  - powering-on: 正在开启，实例可以绑定、解绑  - powering-off: 正在关闭，实例可以绑定、解绑  - rebooting: 正在重启，实例可以绑定、解绑  - delete_wait: 等待删除，集群与实例不允许任何操作  - NO_TASK: 不展示
	Task string `json:"task"`

	// 实例的当前版本
	Version string `json:"version"`

	// 虚拟私有云
	VpcId string `json:"vpc_id"`

	// 可用区
	Zone string `json:"zone"`
}

func (o AuditInstanceListBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditInstanceListBean struct{}"
	}

	return strings.Join([]string{"AuditInstanceListBean", string(data)}, " ")
}
