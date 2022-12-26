package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CBH实例详情
type InstanceDetail struct {

	// 弹性ip
	Publicip string `json:"publicip"`

	// 过期时间，要求用UTC时间表示
	ExpTime string `json:"expTime"`

	// 开始时间，要求用UTC时间表示
	StartTime string `json:"startTime"`

	// 结束时间，要求用UTC时间表示
	EndTime string `json:"endTime"`

	// 释放时间，要求用UTC时间表示
	ReleaseTime string `json:"releaseTime"`

	// 实例名称
	Name string `json:"name"`

	// 实例的server id
	InstanceId string `json:"instanceId"`

	// 实例私有ip
	PrivateIp string `json:"privateIp"`

	// 实例当前的任务状态
	TaskStatus string `json:"taskStatus"`

	// 实例状态
	Status string `json:"status"`

	// 实例创建时间，要求用UTC时间表示
	Created string `json:"created"`

	// 实例所在region
	Region string `json:"region"`

	// 实例所在可用区id
	Zone string `json:"zone"`

	// 实例所在可用区名称
	AvailabilityZoneDisplay string `json:"availability_zone_display"`

	// 实例所在vpc的id
	VpcId string `json:"vpcId"`

	// 实例所在子网的id
	SubnetId string `json:"subnetId"`

	// 实例所属的安全组的id
	SecurityGroupId string `json:"securityGroupId"`

	// 实例规格
	Specification string `json:"specification"`

	// 实例镜像是否可以升级
	Update string `json:"update"`

	// 在创建实例过程中的过程状态信息
	CreateinstanceStatus string `json:"createinstanceStatus"`

	// 创建实例失败原因
	FailReason string `json:"failReason"`

	// 实例ID
	InstanceKey string `json:"instanceKey"`

	// 订单id
	OrderId string `json:"orderId"`

	// 租户购买的时长
	PeriodNum string `json:"periodNum"`

	// 实例的资源id
	ResourceId string `json:"resourceId"`

	// 堡垒机类型
	BastionType string `json:"bastion_type"`

	// 前端是否显示扩容按钮
	AlterPermit string `json:"alterPermit"`

	// 实例绑定的弹性IP的id
	PublicId string `json:"publicId"`

	// 实例镜像当前版本号
	BastionVersion string `json:"bastionVersion"`

	// 实例镜像最新版本号
	NewBastionVersion string `json:"newBastionVersion"`

	// 实例状态
	InstanceStatus string `json:"instanceStatus"`

	// 实例描述
	InstanceDescription string `json:"instanceDescription"`

	// 是否支持续费
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 备可用分区
	SlaveZone *string `json:"slaveZone,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	// 实例类型
	InstanceType *string `json:"instanceType,omitempty"`

	// 主备ID
	HaId *string `json:"haId,omitempty"`

	// 备可用分区名称
	SlaveZoneDisplay *string `json:"slaveZoneDisplay,omitempty"`

	// 端口号
	WebPort *string `json:"webPort,omitempty"`

	// 浮动ip
	Vip *string `json:"vip,omitempty"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}
