package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseAuditV2 struct {

	// 绑定数据库类型
	BindingDbType *string `json:"binding_db_type,omitempty"`

	// 付费模式  - Period：包周期 - Demand：按需。
	ChargeModel string `json:"charge_model"`

	// 备注
	Comment string `json:"comment"`

	// 已配置数据库数量
	ConfigNum int32 `json:"config_num"`

	// IPV6
	ConnectIpv6 *string `json:"connectIpv6,omitempty"`

	// IPV4
	ConnectIp string `json:"connect_ip"`

	// CPU数量
	Cpu int32 `json:"cpu"`

	// 创建时间
	Created string `json:"created"`

	// 数据库数量限额
	DatabaseLimit int32 `json:"database_limit"`

	// 实例结果状态 - 1：冻结可释放  - 2：冻结不可释放 - 3：冻结后不可续费
	Effect int32 `json:"effect"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 过期时间
	Expired string `json:"expired"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 在线天数
	KeepDays string `json:"keep_days"`

	// 实例名称
	Name string `json:"name"`

	// 最新版本
	NewVersion string `json:"new_version"`

	// 端口ID
	PortId string `json:"port_id"`

	// 内存大小
	Ram int32 `json:"ram"`

	// 所属区域
	Region string `json:"region"`

	// 剩余天数
	RemainDays string `json:"remain_days"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源规格编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 冻结场景  - POLICE: 公安冻结  - ILLEGAL: 违规冻结  - VERIFY: 未实名认证冻结  - PARTNER: 合作伙伴冻结 - ARREARS: 普通冻结（普通）
	Scene string `json:"scene"`

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`

	// 规格
	Specification string `json:"specification"`

	// 实例状态  - SHUTOFF: 已关闭  - ACTIVE: 运行中，允许任何操作   - DELETING: 删除中，不允许任何操作  - BUILD: 创建中，不允许任何操作  - DELETED: 已删除，不需要展示  - ERROR: 故障，只允许删除  - HAWAIT: 等待备机创建成功，不允许任何操作  - FROZEN: 已冻结，只允许续费、绑定/解绑  - UPGRADING: 升级中，不允许升级操作
	Status string `json:"status"`

	// 子网ID
	SubnetId string `json:"subnet_id"`

	// 功能列表
	SupportedFeature *[]string `json:"supported_feature,omitempty"`

	// 任务状态  - powering-on: 正在开启，实例可以绑定、解绑  - powering-off: 正在关闭，实例可以绑定、解绑  - rebooting: 正在重启，实例可以绑定、解绑  - delete_wait: 等待删除，集群与实例不允许任何操作  - NO_TASK: 不展示
	Task string `json:"task"`

	// 时区
	Timezone *string `json:"timezone,omitempty"`

	// 升级日志
	UpgradeLog *string `json:"upgrade_log,omitempty"`

	// 实例版本
	Version string `json:"version"`

	// VPC私有云ID
	VpcId string `json:"vpc_id"`

	// 可用区
	Zone string `json:"zone"`
}

func (o ResponseAuditV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseAuditV2 struct{}"
	}

	return strings.Join([]string{"ResponseAuditV2", string(data)}, " ")
}
