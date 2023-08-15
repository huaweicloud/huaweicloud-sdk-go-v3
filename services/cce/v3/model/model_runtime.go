package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Runtime struct {

	// 容器运行时, 默认场景： - 1.25以下集群：默认为\"docker\" - 1.25及以上集群，随操作系统变化，默认的容器运行时不同：操作系统为欧拉2.5、欧拉2.8的节点默认为\"docker\"，其余操作系统的节点默认为\"containerd\"
	Name *RuntimeName `json:"name,omitempty"`
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
