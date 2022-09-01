package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RomaInstanceCheckListRespInstances struct {

	// 实例ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 实例名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 实例描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 实例规格ID
	FlavorId *string `json:"flavor_id,omitempty" xml:"flavor_id"`

	// 实例规格类型
	FlavorType *string `json:"flavor_type,omitempty" xml:"flavor_type"`

	// CPU架构类型，取值如下： - x86_64: x86架构 - aarch64: arm架构
	CpuArch *RomaInstanceCheckListRespInstancesCpuArch `json:"cpu_arch,omitempty" xml:"cpu_arch"`

	// 实例指定虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 实例指定虚拟私有云子网ID
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 实例指定安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 是否开启公网访问，开启时publicip_id字段必填。
	PublicipEnable *bool `json:"publicip_enable,omitempty" xml:"publicip_enable"`

	// 实例绑定的弹性公网地址ID
	PublicipId *string `json:"publicip_id,omitempty" xml:"publicip_id"`

	// 实例绑定的弹性公网地址
	PublicipAddress *string `json:"publicip_address,omitempty" xml:"publicip_address"`

	// 实例运行状态
	Status *RomaInstanceCheckListRespInstancesStatus `json:"status,omitempty" xml:"status"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误消息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 实例计费模式
	ChargeType *RomaInstanceCheckListRespInstancesChargeType `json:"charge_type,omitempty" xml:"charge_type"`

	// 租户项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 运维开始时间
	MaintainBegin *string `json:"maintain_begin,omitempty" xml:"maintain_begin"`

	// 运维结束时间
	MaintainEnd *string `json:"maintain_end,omitempty" xml:"maintain_end"`

	// 创实例使用的可用区列表
	AvailableZoneIds *[]string `json:"available_zone_ids,omitempty" xml:"available_zone_ids"`

	// 实例所属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o RomaInstanceCheckListRespInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RomaInstanceCheckListRespInstances struct{}"
	}

	return strings.Join([]string{"RomaInstanceCheckListRespInstances", string(data)}, " ")
}

type RomaInstanceCheckListRespInstancesCpuArch struct {
	value string
}

type RomaInstanceCheckListRespInstancesCpuArchEnum struct {
	X86_64  RomaInstanceCheckListRespInstancesCpuArch
	AARCH64 RomaInstanceCheckListRespInstancesCpuArch
}

func GetRomaInstanceCheckListRespInstancesCpuArchEnum() RomaInstanceCheckListRespInstancesCpuArchEnum {
	return RomaInstanceCheckListRespInstancesCpuArchEnum{
		X86_64: RomaInstanceCheckListRespInstancesCpuArch{
			value: "x86_64",
		},
		AARCH64: RomaInstanceCheckListRespInstancesCpuArch{
			value: "aarch64",
		},
	}
}

func (c RomaInstanceCheckListRespInstancesCpuArch) Value() string {
	return c.value
}

func (c RomaInstanceCheckListRespInstancesCpuArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RomaInstanceCheckListRespInstancesCpuArch) UnmarshalJSON(b []byte) error {
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

type RomaInstanceCheckListRespInstancesStatus struct {
	value string
}

type RomaInstanceCheckListRespInstancesStatusEnum struct {
	CREATING      RomaInstanceCheckListRespInstancesStatus
	CREATE_FAILED RomaInstanceCheckListRespInstancesStatus
	ERROR         RomaInstanceCheckListRespInstancesStatus
	RUNNING       RomaInstanceCheckListRespInstancesStatus
	FREEZING      RomaInstanceCheckListRespInstancesStatus
	FROZEN        RomaInstanceCheckListRespInstancesStatus
	DELETING      RomaInstanceCheckListRespInstancesStatus
	DELETE_FALIED RomaInstanceCheckListRespInstancesStatus
}

func GetRomaInstanceCheckListRespInstancesStatusEnum() RomaInstanceCheckListRespInstancesStatusEnum {
	return RomaInstanceCheckListRespInstancesStatusEnum{
		CREATING: RomaInstanceCheckListRespInstancesStatus{
			value: "CREATING",
		},
		CREATE_FAILED: RomaInstanceCheckListRespInstancesStatus{
			value: "CREATE_FAILED",
		},
		ERROR: RomaInstanceCheckListRespInstancesStatus{
			value: "ERROR",
		},
		RUNNING: RomaInstanceCheckListRespInstancesStatus{
			value: "RUNNING",
		},
		FREEZING: RomaInstanceCheckListRespInstancesStatus{
			value: "FREEZING",
		},
		FROZEN: RomaInstanceCheckListRespInstancesStatus{
			value: "FROZEN",
		},
		DELETING: RomaInstanceCheckListRespInstancesStatus{
			value: "DELETING",
		},
		DELETE_FALIED: RomaInstanceCheckListRespInstancesStatus{
			value: "DELETE_FALIED",
		},
	}
}

func (c RomaInstanceCheckListRespInstancesStatus) Value() string {
	return c.value
}

func (c RomaInstanceCheckListRespInstancesStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RomaInstanceCheckListRespInstancesStatus) UnmarshalJSON(b []byte) error {
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

type RomaInstanceCheckListRespInstancesChargeType struct {
	value string
}

type RomaInstanceCheckListRespInstancesChargeTypeEnum struct {
	PRE_PAID  RomaInstanceCheckListRespInstancesChargeType
	POST_PAID RomaInstanceCheckListRespInstancesChargeType
}

func GetRomaInstanceCheckListRespInstancesChargeTypeEnum() RomaInstanceCheckListRespInstancesChargeTypeEnum {
	return RomaInstanceCheckListRespInstancesChargeTypeEnum{
		PRE_PAID: RomaInstanceCheckListRespInstancesChargeType{
			value: "prePaid",
		},
		POST_PAID: RomaInstanceCheckListRespInstancesChargeType{
			value: "postPaid",
		},
	}
}

func (c RomaInstanceCheckListRespInstancesChargeType) Value() string {
	return c.value
}

func (c RomaInstanceCheckListRespInstancesChargeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RomaInstanceCheckListRespInstancesChargeType) UnmarshalJSON(b []byte) error {
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
