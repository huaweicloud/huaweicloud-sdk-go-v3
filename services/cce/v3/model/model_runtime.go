package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Runtime struct {

	// 容器运行时，默认场景： - v1.25以下集群：默认为\"docker\" - v1.25及以上集群，随操作系统变化，默认的容器运行时不同：操作系统为EulerOS 2.5[、EulerOS 2.8](tag:hws,hws_hk)的节点默认为\"docker\"，其余操作系统的节点默认为\"containerd\"
	Name *RuntimeName `json:"name,omitempty"`

	// **参数解释**： 容器运行时子类别。 **约束限制**： 仅CCE Turbo集群下弹性云服务器-物理机类型节点且上级运行时为containerd场景支持使用安全运行时。 **取值范围**： - runc: 普通运行时。 - kata: 安全运行时，需配套c6、c7系列弹性云服务器-物理机，支持的操作系统为EulerOS 2.10。 - kuasar-vmm: 安全运行时v2，支持kc2、ki2、c7、ac8h系列弹性服务器-物理机，配套操作系统为HCE 2.0，集群版本需为v1.28.15-r70、v1.29.15-r30、v1.30.14-r30、v1.31.10-r30、v1.32.6-r30、v1.33.5-r20或以上版本。  **默认取值**： runc
	RuntimeClass *RuntimeRuntimeClass `json:"runtimeClass,omitempty"`
}

func (o Runtime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Runtime struct{}"
	}

	return strings.Join([]string{"Runtime", string(data)}, " ")
}

type RuntimeName struct {
	value string
}

type RuntimeNameEnum struct {
	DOCKER     RuntimeName
	CONTAINERD RuntimeName
}

func GetRuntimeNameEnum() RuntimeNameEnum {
	return RuntimeNameEnum{
		DOCKER: RuntimeName{
			value: "docker",
		},
		CONTAINERD: RuntimeName{
			value: "containerd",
		},
	}
}

func (c RuntimeName) Value() string {
	return c.value
}

func (c RuntimeName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuntimeName) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type RuntimeRuntimeClass struct {
	value string
}

type RuntimeRuntimeClassEnum struct {
	RUNC       RuntimeRuntimeClass
	KATA       RuntimeRuntimeClass
	KUASAR_VMM RuntimeRuntimeClass
}

func GetRuntimeRuntimeClassEnum() RuntimeRuntimeClassEnum {
	return RuntimeRuntimeClassEnum{
		RUNC: RuntimeRuntimeClass{
			value: "runc",
		},
		KATA: RuntimeRuntimeClass{
			value: "kata",
		},
		KUASAR_VMM: RuntimeRuntimeClass{
			value: "kuasar-vmm",
		},
	}
}

func (c RuntimeRuntimeClass) Value() string {
	return c.value
}

func (c RuntimeRuntimeClass) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuntimeRuntimeClass) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
