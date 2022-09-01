package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CBH实例详情
type InstanceDetail struct {

	// 弹性ip
	Publicip string `json:"publicip" xml:"publicip"`

	// 过期时间
	ExpTime *string `json:"exp_time,omitempty" xml:"exp_time"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 释放时间
	ReleaseTime *string `json:"release_time,omitempty" xml:"release_time"`

	// 实例名称
	Name string `json:"name" xml:"name"`

	// 实例的server id
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 实例私有ip
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip"`

	// 实例当前的任务状态
	TaskStatus *string `json:"task_status,omitempty" xml:"task_status"`

	// 实例状态
	Status string `json:"status" xml:"status"`

	// 实例创建时间
	Created string `json:"created" xml:"created"`

	// 实例所在region
	Region string `json:"region" xml:"region"`

	// 实例所在可用区id
	Zone string `json:"zone" xml:"zone"`

	// 实例所在可用区名称
	AvailabilityZoneDisplay string `json:"availability_zone_display" xml:"availability_zone_display"`

	// 实例所在vpc的id
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 实例所在子网的id
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 实例所属的安全组的id
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 实例规格
	Specification string `json:"specification" xml:"specification"`

	// 实例镜像是否可以升级
	Update string `json:"update" xml:"update"`

	// 在创建实例过程中的过程状态信息
	CreateinstanceStatus *string `json:"createinstance_status,omitempty" xml:"createinstance_status"`

	// 创建实例失败原因
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason"`

	// 实例ID
	InstanceKey *string `json:"instance_key,omitempty" xml:"instance_key"`

	// 订单id
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 租户购买的时长
	PeriodNum *string `json:"period_num,omitempty" xml:"period_num"`

	// 实例的资源id
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 堡垒机类型
	BastionType string `json:"bastion_type" xml:"bastion_type"`

	// 实例绑定的弹性IP的id
	PublicId *string `json:"public_id,omitempty" xml:"public_id"`

	// 前端是否显示扩容按钮
	AlterPermit *string `json:"alter_permit,omitempty" xml:"alter_permit"`

	// 实例镜像当前版本号
	BastionVersion *string `json:"bastion_version,omitempty" xml:"bastion_version"`

	// 实例镜像最新版本号
	NewBastionVersion *string `json:"new_bastion_version,omitempty" xml:"new_bastion_version"`

	// 实例状态
	InstanceStatus *string `json:"instance_status,omitempty" xml:"instance_status"`

	// 实例描述
	InstanceDescription *string `json:"instance_description,omitempty" xml:"instance_description"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}
