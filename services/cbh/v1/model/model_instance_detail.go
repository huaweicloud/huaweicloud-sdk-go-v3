package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CBH实例详情
type InstanceDetail struct {

	// 云堡垒机实例弹性公网IP，返回默认值null
	Publicip string `json:"publicip"`

	// 云堡垒机实例过期时间。
	ExpTime string `json:"expTime"`

	// 云堡垒机实例开始时间，使用时间戳格式表示。
	StartTime string `json:"startTime"`

	// 云堡垒机实例结束时间，使用时间戳格式表示。
	EndTime string `json:"endTime"`

	// 云堡垒机释放时间，使用时间戳格式表示。
	ReleaseTime string `json:"releaseTime"`

	// 云堡垒机实例名称。
	Name string `json:"name"`

	// 云堡垒机实例ID，UUID格式。
	InstanceId string `json:"instanceId"`

	// 云堡垒机实例私有ip。
	PrivateIp string `json:"privateIp"`

	// 云堡垒机实例当前的任务状态。 - powering-on 开启 - powering-off 关闭 - rebooting 重启 - delete_wait 删除 - frozen 冻结 - NO_TASK 运行 - unfrozen 解冻 - alter 变更 - updating 升级中 - configuring-ha 配置HA
	TaskStatus string `json:"taskStatus"`

	// 云堡垒机实例状态。 - SHUTOFF 已关闭 - ACTIVE 运行中 - DELETING 删除中 - BUILD 创建中 - DELETED 已删除 - ERROR 故障 - HAWAIT 等待备机创建成功 - FROZEN 已冻结 - UPGRADING 升级中 - UNPAID 待支付 - RESIZE 规格变更中 - DILATATION 扩容中 - HA 配置HA中
	Status string `json:"status"`

	// 云堡垒机实例创建时间，使用UTC时间表示。
	Created string `json:"created"`

	// 云堡垒机实例所在局点。
	Region string `json:"region"`

	// 云堡垒机实例所在可用区。
	Zone string `json:"zone"`

	// 云堡垒机实例所在可用区中文名称。
	AvailabilityZoneDisplay string `json:"availability_zone_display"`

	// 云堡垒机实例所在虚拟私有云的VPC ID。
	VpcId string `json:"vpcId"`

	// 云堡垒机实例所在子网的ID。
	SubnetId string `json:"subnetId"`

	// 云堡垒机实例所属的安全组的ID。
	SecurityGroupId string `json:"securityGroupId"`

	// 云堡垒机实例规格。
	Specification string `json:"specification"`

	// 云堡垒机实例是否可以升级。 - NEW，可以升级 - OLD，不能升级
	Update string `json:"update"`

	// 云堡垒机实例在创建实例过程中的过程状态信息。 - Waiting for payment，等待支付 - creating-network，创建网络 - creating-server，创建服务 - tranfering-horizontal-network，网络打通 - adding-policy-route，添加路由策略 - configing-dns，配置DNS - starting-cbs-service，服务运行中 - setting-init-conf，初始化 - buying-EIP，购买弹性公网IP
	CreateinstanceStatus string `json:"createinstanceStatus"`

	// 云堡垒机实例创建实例失败原因。
	FailReason string `json:"failReason"`

	// 云堡垒机实例key。
	InstanceKey string `json:"instanceKey"`

	// 订单ID。
	OrderId string `json:"orderId"`

	// 云堡垒机实例订购周期数。
	PeriodNum string `json:"periodNum"`

	// 云堡垒机实例的资源id,UUID格式显示。
	ResourceId string `json:"resourceId"`

	// 云堡垒机实例堡垒机类型。 - OEM
	BastionType string `json:"bastion_type"`

	// 云堡垒机实例是否可以扩容。 - 1 开启扩容 - 0 关闭扩容
	AlterPermit string `json:"alterPermit"`

	// 云堡垒机实例绑定公网的弹性IP的ID，UUID格式表示。
	PublicId string `json:"publicId"`

	// 云堡垒机实例当前版本。
	BastionVersion string `json:"bastionVersion"`

	// 云堡垒机实例可以升级的版本。
	NewBastionVersion string `json:"newBastionVersion"`

	// 云堡垒机实例状态。 - building  创建中 - deleting  删除中 - deleted 删除了 - unpaid  未支付 - upgrading 升级中 - resizing  扩容中 - abnormal  异常 - error 故障 - ok  正常
	InstanceStatus string `json:"instanceStatus"`

	// 云堡垒机实例状态描述。
	InstanceDescription string `json:"instanceDescription"`

	// 备可用分区，默认返回null。
	SlaveZone *string `json:"slaveZone,omitempty"`

	// 云堡垒机实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	// 云堡垒机实例类型。 - null  单机默认返回null - master  HA时返回主机 - slave HA时返回备机
	InstanceType *string `json:"instanceType,omitempty"`

	// 云堡垒机实例主备ID。 - 单机堡垒机实例默认返回null - HA堡垒机实例返回主机HAID
	HaId *string `json:"haId,omitempty"`

	// 云堡垒机实例备机可用分区名称。 单机堡垒机实例和备机堡垒机实例返回null，HA堡垒机实例主机返回备机所在可用区名称。
	SlaveZoneDisplay *string `json:"slaveZoneDisplay,omitempty"`

	// 云堡垒机实例WEB界面访问的端口号。
	WebPort *string `json:"webPort,omitempty"`

	// 云堡垒机实例浮动ip。返回默认值null
	Vip *string `json:"vip,omitempty"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}
