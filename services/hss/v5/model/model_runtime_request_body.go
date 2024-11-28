package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RuntimeRequestBody struct {

	// 运行时名称：   - crio_endpoint：CRIO   - containerd_endpoint：Containerd   - docker_endpoint：Docker   - isulad_endpoint：Isulad   - podman_endpoint：Podman
	RuntimeName *RuntimeRequestBodyRuntimeName `json:"runtime_name,omitempty"`

	// 运行时路径
	RuntimePath *string `json:"runtime_path,omitempty"`
}

func (o RuntimeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeRequestBody struct{}"
	}

	return strings.Join([]string{"RuntimeRequestBody", string(data)}, " ")
}

type RuntimeRequestBodyRuntimeName struct {
	value string
}

type RuntimeRequestBodyRuntimeNameEnum struct {
	CRIO_ENDPOINT       RuntimeRequestBodyRuntimeName
	CONTAINERD_ENDPOINT RuntimeRequestBodyRuntimeName
	DOCKER_ENDPOINT     RuntimeRequestBodyRuntimeName
	ISULAD_ENDPOINT     RuntimeRequestBodyRuntimeName
	PODMAN_ENDPOINT     RuntimeRequestBodyRuntimeName
}

func GetRuntimeRequestBodyRuntimeNameEnum() RuntimeRequestBodyRuntimeNameEnum {
	return RuntimeRequestBodyRuntimeNameEnum{
		CRIO_ENDPOINT: RuntimeRequestBodyRuntimeName{
			value: "crio_endpoint",
		},
		CONTAINERD_ENDPOINT: RuntimeRequestBodyRuntimeName{
			value: "containerd_endpoint",
		},
		DOCKER_ENDPOINT: RuntimeRequestBodyRuntimeName{
			value: "docker_endpoint",
		},
		ISULAD_ENDPOINT: RuntimeRequestBodyRuntimeName{
			value: "isulad_endpoint",
		},
		PODMAN_ENDPOINT: RuntimeRequestBodyRuntimeName{
			value: "podman_endpoint",
		},
	}
}

func (c RuntimeRequestBodyRuntimeName) Value() string {
	return c.value
}

func (c RuntimeRequestBodyRuntimeName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuntimeRequestBodyRuntimeName) UnmarshalJSON(b []byte) error {
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
