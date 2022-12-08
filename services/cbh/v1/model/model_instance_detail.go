package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CBH实例详情
type InstanceDetail struct {

	// 弹性ip
	PublicIp string `json:"public_ip"`

	// 过期时间，要求用UTC时间表示
	ExpTime string `json:"exp_time"`

	// 开始时间，要求用UTC时间表示
	StartTime string `json:"start_time"`

	// 结束时间，要求用UTC时间表示
	EndTime string `json:"end_time"`

	// 释放时间，要求用UTC时间表示
	ReleaseTime string `json:"release_time"`

	// 实例名称
	Name string `json:"name"`

	// 实例的server id
	InstanceId string `json:"instance_id"`

	// 实例私有ip
	PrivateIp string `json:"private_ip"`

	// 实例当前的任务状态
	TaskStatus string `json:"task_status"`

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
	VpcId string `json:"vpc_id"`

	// 实例所在子网的id
	SubnetId string `json:"subnet_id"`

	// 实例所属的安全组的id
	SecurityGroupId string `json:"security_group_id"`

	// 实例规格
	Specification string `json:"specification"`

	// 实例镜像是否可以升级
	Update string `json:"update"`

	// 在创建实例过程中的过程状态信息
	CreateinstanceStatus string `json:"createinstance_status"`

	// 创建实例失败原因
	FailReason string `json:"fail_reason"`

	// 实例ID
	InstanceKey string `json:"instance_key"`

	// 订单id
	OrderId string `json:"order_id"`

	// 租户购买的时长
	PeriodNum string `json:"period_num"`

	// 实例的资源id
	ResourceId string `json:"resource_id"`

	// 堡垒机类型
	BastionType string `json:"bastion_type"`

	// 前端是否显示扩容按钮
	AlterPermit string `json:"alter_permit"`

	// 实例绑定的弹性IP的id
	PublicId string `json:"public_id"`

	// 实例镜像当前版本号
	BastionVersion string `json:"bastion_version"`

	// 实例镜像最新版本号
	NewBastionVersion string `json:"new_bastion_version"`

	// 实例状态
	InstanceStatus string `json:"instance_status"`

	// 实例描述
	InstanceDescription string `json:"instance_description"`

	// 是否支持续费
	IsAutoRenew int32 `json:"is_auto_renew"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}
