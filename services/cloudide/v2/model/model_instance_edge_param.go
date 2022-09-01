package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceEdgeParam struct {

	// cpu架构 x86|arm
	Arch *InstanceEdgeParamArch `json:"arch,omitempty" xml:"arch"`

	// cpu规格.arm架构支持4U8G，x86架构支持1U1G,2U4G,2U8G 与技术栈配置的规格对应，可通过技术栈管理ListStacks接口获取。如果标签不为空，以标签配置的技术栈规格为准。 quantum技术栈，x86架构cpu规格为2U8G;其他技术栈，x86架构cpu规格为1U1G,2U4G
	CpuMemory InstanceEdgeParamCpuMemory `json:"cpu_memory" xml:"cpu_memory"`

	// 描述。长度不操过100个字符
	Description *string `json:"description,omitempty" xml:"description"`

	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间
	InstanceName string `json:"instance_name" xml:"instance_name"`

	// 租户id（对应华为云帐号的domainId）
	InstanceUserDomainId *string `json:"instance_user_domain_id,omitempty" xml:"instance_user_domain_id"`

	// 用户id
	InstanceUserId *string `json:"instance_user_id,omitempty" xml:"instance_user_id"`

	// 是否页面显示（以标签配置为准）
	IsTemporary *bool `json:"is_temporary,omitempty" xml:"is_temporary"`

	// 插件列表
	Plugins *[]Plugin `json:"plugins,omitempty" xml:"plugins"`

	// PVC规格 5GB|10GB|20GB
	PvcQuantity InstanceEdgeParamPvcQuantity `json:"pvc_quantity" xml:"pvc_quantity"`

	// 自动休眠时长。 arm架构,自动休眠时长只能设置成30，60。x86架构可取值为30，60，240，1440和-1。除-1外，其它值的单位为“分钟”。实例无操作超过自动休眠时长后，将会被暂停（已保存的数据不会被删除）。-1表示实例不会自动停止。
	RefreshTime *string `json:"refresh_time,omitempty" xml:"refresh_time"`

	// 技术栈ID，通过技术栈管理ListStacks接口获取。
	StackId string `json:"stack_id" xml:"stack_id"`
}

func (o InstanceEdgeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceEdgeParam struct{}"
	}

	return strings.Join([]string{"InstanceEdgeParam", string(data)}, " ")
}

type InstanceEdgeParamArch struct {
	value string
}

type InstanceEdgeParamArchEnum struct {
	X86 InstanceEdgeParamArch
	ARM InstanceEdgeParamArch
}

func GetInstanceEdgeParamArchEnum() InstanceEdgeParamArchEnum {
	return InstanceEdgeParamArchEnum{
		X86: InstanceEdgeParamArch{
			value: "x86",
		},
		ARM: InstanceEdgeParamArch{
			value: "arm",
		},
	}
}

func (c InstanceEdgeParamArch) Value() string {
	return c.value
}

func (c InstanceEdgeParamArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceEdgeParamArch) UnmarshalJSON(b []byte) error {
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

type InstanceEdgeParamCpuMemory struct {
	value string
}

type InstanceEdgeParamCpuMemoryEnum struct {
	E_1_U1_G InstanceEdgeParamCpuMemory
	E_2_U4_G InstanceEdgeParamCpuMemory
	E_2_U8_G InstanceEdgeParamCpuMemory
	E_4_U8_G InstanceEdgeParamCpuMemory
}

func GetInstanceEdgeParamCpuMemoryEnum() InstanceEdgeParamCpuMemoryEnum {
	return InstanceEdgeParamCpuMemoryEnum{
		E_1_U1_G: InstanceEdgeParamCpuMemory{
			value: "1U1G",
		},
		E_2_U4_G: InstanceEdgeParamCpuMemory{
			value: "2U4G",
		},
		E_2_U8_G: InstanceEdgeParamCpuMemory{
			value: "2U8G",
		},
		E_4_U8_G: InstanceEdgeParamCpuMemory{
			value: "4U8G",
		},
	}
}

func (c InstanceEdgeParamCpuMemory) Value() string {
	return c.value
}

func (c InstanceEdgeParamCpuMemory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceEdgeParamCpuMemory) UnmarshalJSON(b []byte) error {
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

type InstanceEdgeParamPvcQuantity struct {
	value string
}

type InstanceEdgeParamPvcQuantityEnum struct {
	E_5_GB  InstanceEdgeParamPvcQuantity
	E_10_GB InstanceEdgeParamPvcQuantity
	E_20_GB InstanceEdgeParamPvcQuantity
}

func GetInstanceEdgeParamPvcQuantityEnum() InstanceEdgeParamPvcQuantityEnum {
	return InstanceEdgeParamPvcQuantityEnum{
		E_5_GB: InstanceEdgeParamPvcQuantity{
			value: "5GB",
		},
		E_10_GB: InstanceEdgeParamPvcQuantity{
			value: "10GB",
		},
		E_20_GB: InstanceEdgeParamPvcQuantity{
			value: "20GB",
		},
	}
}

func (c InstanceEdgeParamPvcQuantity) Value() string {
	return c.value
}

func (c InstanceEdgeParamPvcQuantity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceEdgeParamPvcQuantity) UnmarshalJSON(b []byte) error {
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
