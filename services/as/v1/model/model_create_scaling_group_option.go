package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建伸缩组请求
type CreateScalingGroupOption struct {
	// 伸缩组名称(1-64个字符)，只能包含中文、字母、数字、下划线、中划线。

	ScalingGroupName string `json:"scaling_group_name"`
	// 伸缩配置ID，通过查询弹性伸缩配置列表接口获取。

	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`
	// 期望实例数量，默认值为最小实例数。最小实例数<=期望实例数<=最大实例数

	DesireInstanceNumber *int32 `json:"desire_instance_number,omitempty"`
	// 最小实例数量，默认值为0。

	MinInstanceNumber *int32 `json:"min_instance_number,omitempty"`
	// 最大实例数量，默认值为0。

	MaxInstanceNumber *int32 `json:"max_instance_number,omitempty"`
	// 冷却时间，取值范围0-86400，默认为900，单位是秒。 只针对告警策略生效，定时、周期策略和手动触发策略不受该参数限制。

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`
	// 弹性负载均衡（经典型）监听器ID，最多支持绑定3个负载均衡监听器，多个负载均衡监听器ID以逗号分隔。首先使用vpc_id通过查询ELB服务负载均衡器列表接口获取负载均衡器的ID，详见《弹性负载均衡API参考》的“查询负载均衡器列表”，再使用该ID查询监听器列表获取，详见《弹性负载均衡API参考》的“查询监听器列表”。

	LbListenerId *string `json:"lb_listener_id,omitempty"`
	// 弹性负载均衡器（增强型）信息，最多支持绑定3个负载均衡。该字段与lb_listener_id互斥。

	LbaasListeners *[]LbaasListeners `json:"lbaas_listeners,omitempty"`
	// 可用分区信息。弹性伸缩活动中自动添加的云服务器会被创建在指定的可用区中。如果没有指定可用分区，会由系统自动指定可用分区。详情请参考地区和终端节点。

	AvailableZones *[]string `json:"available_zones,omitempty"`
	// 网络信息，最多支持选择5个子网，传入的第一个子网默认作为云服务器的主网卡。获取子网信息请参考[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。

	Networks []Networks `json:"networks"`
	// 安全组信息，最多支持选择1个安全组。使用vpc_id通过查询VPC服务安全组列表接口获取，详见《虚拟私有云API参考》的“查询安全组列表”。当伸缩配置和伸缩组同时指定安全组时，将以伸缩配置中的安全组为准；当伸缩配置和伸缩组都没有指定安全组时，将使用默认安全组。为了使用灵活性更高，推荐在伸缩配置中指定安全组。

	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`
	// VPC信息，通过查询VPC服务VPC列表接口获取，详见《虚拟私有云API参考》的“查询VPC列表”。

	VpcId string `json:"vpc_id"`
	// 伸缩组实例健康检查方式：ELB_AUDIT和NOVA_AUDIT。当伸缩组参数中设置负载均衡时，默认为ELB_AUDIT；否则默认为NOVA_AUDIT。ELB_AUDIT表示负载均衡健康检查方式，在有监听器的伸缩组中有效。NOVA_AUDIT表示弹性伸缩自带的健康检查方式。

	HealthPeriodicAuditMethod *CreateScalingGroupOptionHealthPeriodicAuditMethod `json:"health_periodic_audit_method,omitempty"`
	// 伸缩组实例的健康检查周期，可设置为1、5、15、60、180（分钟），若不设置该参数，默认为5。若设置为0，可以实现10秒级健康检查。

	HealthPeriodicAuditTime *int32 `json:"health_periodic_audit_time,omitempty"`
	// 伸缩组实例健康状况检查宽限期，取值范围0-86400，单位是秒。当实例加入伸缩组并且进入已启用状态后，健康状况检查宽限期才会启动，伸缩组会等健康状况检查宽限期结束后才检查实例的运行状况。当伸缩组实例健康检查方式为ELB_AUDIT时，该参数生效，若不设置该参数，默认为600秒。

	HealthPeriodicAuditGracePeriod *int32 `json:"health_periodic_audit_grace_period,omitempty"`
	// 伸缩组实例移除策略：OLD_CONFIG_OLD_INSTANCE（默认）：从根据“较早创建的配置”创建的实例中筛选出较早创建的实例被优先移除。OLD_CONFIG_NEW_INSTANCE：从根据“较早创建的配置”创建的实例中筛选出较新创建的实例被优先移除。OLD_INSTANCE：较早创建的实例被优先移除。NEW_INSTANCE：较新创建的实例将被优先移除。

	InstanceTerminatePolicy *CreateScalingGroupOptionInstanceTerminatePolicy `json:"instance_terminate_policy,omitempty"`
	// 通知方式：EMAIL为发送邮件通知。该通知方式即将被废除，建议给弹性伸缩组配置通知功能。详见通知。

	Notifications *[]string `json:"notifications,omitempty"`
	// 配置删除云服务器时是否删除云服务器绑定的弹性IP。取值为true或false，默认为false。true：删除云服务器时，会同时删除绑定在云服务器上的弹性IP。当弹性IP的计费方式为包年包月时，不会被删除。false：删除云服务器时，仅解绑定在云服务器上的弹性IP，不删除弹性IP。

	DeletePublicip *bool `json:"delete_publicip,omitempty"`
	// 配置删除云服务器时是否删除云服务器绑定的数据盘。取值为true或false，默认为false。 true：删除云服务器时，会同时删除绑定在云服务器上的数据盘。当数据盘的计费方式为包年包月时，不会被删除。 false：删除云服务器时，仅解绑定在云服务器上的数据盘，不删除数据盘。

	DeleteVolume *bool `json:"delete_volume,omitempty"`
	// 企业项目ID，用于指定伸缩组归属的企业项目。当伸缩组配置企业项目时，由该伸缩组创建的弹性云服务器将归属于该企业项目。当没有指定企业项目时，将使用企业项目ID为0的默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 伸缩组扩缩容时目标AZ选择的优先级策略：EQUILIBRIUM_DISTRIBUTE（默认）：均衡分布，云服务器扩缩容时优先保证available_zones列表中各AZ下虚拟机数量均衡，当无法在目标AZ下完成虚拟机扩容时，按照PICK_FIRST原则选择其他可用AZ。PICK_FIRST：选择优先，虚拟机扩缩容时目标AZ的选择按照available_zones列表的顺序进行优先级排序。

	MultiAzPriorityPolicy *CreateScalingGroupOptionMultiAzPriorityPolicy `json:"multi_az_priority_policy,omitempty"`
	// 伸缩组描述信息(0-256个字符)

	Description *string `json:"description,omitempty"`
}

func (o CreateScalingGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupOption struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupOption", string(data)}, " ")
}

type CreateScalingGroupOptionHealthPeriodicAuditMethod struct {
	value string
}

type CreateScalingGroupOptionHealthPeriodicAuditMethodEnum struct {
	ELB_AUDIT  CreateScalingGroupOptionHealthPeriodicAuditMethod
	NOVA_AUDIT CreateScalingGroupOptionHealthPeriodicAuditMethod
}

func GetCreateScalingGroupOptionHealthPeriodicAuditMethodEnum() CreateScalingGroupOptionHealthPeriodicAuditMethodEnum {
	return CreateScalingGroupOptionHealthPeriodicAuditMethodEnum{
		ELB_AUDIT: CreateScalingGroupOptionHealthPeriodicAuditMethod{
			value: "ELB_AUDIT",
		},
		NOVA_AUDIT: CreateScalingGroupOptionHealthPeriodicAuditMethod{
			value: "NOVA_AUDIT",
		},
	}
}

func (c CreateScalingGroupOptionHealthPeriodicAuditMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingGroupOptionHealthPeriodicAuditMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateScalingGroupOptionInstanceTerminatePolicy struct {
	value string
}

type CreateScalingGroupOptionInstanceTerminatePolicyEnum struct {
	OLD_CONFIG_OLD_INSTANCE CreateScalingGroupOptionInstanceTerminatePolicy
	OLD_CONFIG_NEW_INSTANCE CreateScalingGroupOptionInstanceTerminatePolicy
	OLD_INSTANCE            CreateScalingGroupOptionInstanceTerminatePolicy
	NEW_INSTANCE            CreateScalingGroupOptionInstanceTerminatePolicy
}

func GetCreateScalingGroupOptionInstanceTerminatePolicyEnum() CreateScalingGroupOptionInstanceTerminatePolicyEnum {
	return CreateScalingGroupOptionInstanceTerminatePolicyEnum{
		OLD_CONFIG_OLD_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_CONFIG_OLD_INSTANCE",
		},
		OLD_CONFIG_NEW_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_CONFIG_NEW_INSTANCE",
		},
		OLD_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_INSTANCE",
		},
		NEW_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "NEW_INSTANCE",
		},
	}
}

func (c CreateScalingGroupOptionInstanceTerminatePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingGroupOptionInstanceTerminatePolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateScalingGroupOptionMultiAzPriorityPolicy struct {
	value string
}

type CreateScalingGroupOptionMultiAzPriorityPolicyEnum struct {
	EQUILIBRIUM_DISTRIBUTE CreateScalingGroupOptionMultiAzPriorityPolicy
	PICK_FIRST             CreateScalingGroupOptionMultiAzPriorityPolicy
}

func GetCreateScalingGroupOptionMultiAzPriorityPolicyEnum() CreateScalingGroupOptionMultiAzPriorityPolicyEnum {
	return CreateScalingGroupOptionMultiAzPriorityPolicyEnum{
		EQUILIBRIUM_DISTRIBUTE: CreateScalingGroupOptionMultiAzPriorityPolicy{
			value: "EQUILIBRIUM_DISTRIBUTE",
		},
		PICK_FIRST: CreateScalingGroupOptionMultiAzPriorityPolicy{
			value: "PICK_FIRST",
		},
	}
}

func (c CreateScalingGroupOptionMultiAzPriorityPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingGroupOptionMultiAzPriorityPolicy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
