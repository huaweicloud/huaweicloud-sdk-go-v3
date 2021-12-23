package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceParam struct {
	// 代理商id，教程活动场景下使用

	AgentId *string `json:"agent_id,omitempty"`
	// cpu架构 x86|arm

	Arch *InstanceParamArch `json:"arch,omitempty"`
	// cpu规格.arm架构支持4U8G，x86架构支持1U1G,2U4G,2U8G 与技术栈配置的规格对应，可通过技术栈管理ListStacksByTag接口获取。如果标签不为空，以标签配置的技术栈规格为准。 quantum技术栈，x86架构cpu规格为2U8G;其他技术栈，x86架构cpu规格为1U1G,2U4G

	CpuMemory InstanceParamCpuMemory `json:"cpu_memory"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间

	DisplayName string `json:"display_name"`
	// 是否页面显示（以标签配置为准）

	IsTemporary *bool `json:"is_temporary,omitempty"`
	// 实例标签（不同的第三方需要和CloudIDE服务共同设定标签），不传默认为default

	LabelTag *string `json:"label_tag,omitempty"`
	// 预装插件列表

	PluginEnableList *[]string `json:"plugin_enable_list,omitempty"`
	// 预装插件参数

	PluginVars map[string]string `json:"plugin_vars,omitempty"`
	// 云服务器对应的portId，小网连接ecs的场景下使用

	PortId *string `json:"port_id,omitempty"`
	// 云服务器ip，小网连接ecs的场景下使用

	PrivateIp *string `json:"private_ip,omitempty"`
	// PVC规格 5GB|10GB|20GB

	PvcQuantity InstanceParamPvcQuantity `json:"pvc_quantity"`
	// 实例的生命周期。 arm架构,生命周期只能设置成30，60。x86架构可取值为30，60，240，1440和-1。除-1外，其它值的单位为“分钟”。实例在到达生命周期后，将会被暂停（已保存的数据不会被删除）。-1表示实例不会自动停止。

	RefreshInterval string `json:"refresh_interval"`
	// 解放号的仓库id，解放号场景下使用

	RepositoryId *int64 `json:"repository_id,omitempty"`
	// 技术栈ID，通过技术栈管理ListStacksByTag接口获取。

	StackId string `json:"stack_id"`
	// 任务类型，教程活动场景下使用

	TaskType *string `json:"task_type,omitempty"`
	// 解放号的token，解放号场景下使用

	Token *string `json:"token,omitempty"`
	// 云服务器对应的vpcId，小网连接ecs的场景下使用

	VpcId *string `json:"vpc_id,omitempty"`
	// 实例授权用户组织名

	InstanceUserDomainName *string `json:"instance_user_domain_name,omitempty"`
	// 实例授权用户名

	InstanceUserName *string `json:"instance_user_name,omitempty"`
}

func (o InstanceParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceParam struct{}"
	}

	return strings.Join([]string{"InstanceParam", string(data)}, " ")
}

type InstanceParamArch struct {
	value string
}

type InstanceParamArchEnum struct {
	X86 InstanceParamArch
	ARM InstanceParamArch
}

func GetInstanceParamArchEnum() InstanceParamArchEnum {
	return InstanceParamArchEnum{
		X86: InstanceParamArch{
			value: "x86",
		},
		ARM: InstanceParamArch{
			value: "arm",
		},
	}
}

func (c InstanceParamArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceParamArch) UnmarshalJSON(b []byte) error {
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

type InstanceParamCpuMemory struct {
	value string
}

type InstanceParamCpuMemoryEnum struct {
	E_1_U1_G InstanceParamCpuMemory
	E_2_U4_G InstanceParamCpuMemory
	E_2_U8_G InstanceParamCpuMemory
	E_4_U8_G InstanceParamCpuMemory
}

func GetInstanceParamCpuMemoryEnum() InstanceParamCpuMemoryEnum {
	return InstanceParamCpuMemoryEnum{
		E_1_U1_G: InstanceParamCpuMemory{
			value: "1U1G",
		},
		E_2_U4_G: InstanceParamCpuMemory{
			value: "2U4G",
		},
		E_2_U8_G: InstanceParamCpuMemory{
			value: "2U8G",
		},
		E_4_U8_G: InstanceParamCpuMemory{
			value: "4U8G",
		},
	}
}

func (c InstanceParamCpuMemory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceParamCpuMemory) UnmarshalJSON(b []byte) error {
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

type InstanceParamPvcQuantity struct {
	value string
}

type InstanceParamPvcQuantityEnum struct {
	E_5_GB  InstanceParamPvcQuantity
	E_10_GB InstanceParamPvcQuantity
	E_20_GB InstanceParamPvcQuantity
}

func GetInstanceParamPvcQuantityEnum() InstanceParamPvcQuantityEnum {
	return InstanceParamPvcQuantityEnum{
		E_5_GB: InstanceParamPvcQuantity{
			value: "5GB",
		},
		E_10_GB: InstanceParamPvcQuantity{
			value: "10GB",
		},
		E_20_GB: InstanceParamPvcQuantity{
			value: "20GB",
		},
	}
}

func (c InstanceParamPvcQuantity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceParamPvcQuantity) UnmarshalJSON(b []byte) error {
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
