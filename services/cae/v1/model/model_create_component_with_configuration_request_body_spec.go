package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateComponentWithConfigurationRequestBodySpec 组件规格信息。
type CreateComponentWithConfigurationRequestBodySpec struct {

	// 语言/运行时。
	Runtime CreateComponentWithConfigurationRequestBodySpecRuntime `json:"runtime"`

	// 实例个数。
	Replica int32 `json:"replica"`

	Build *Build `json:"build,omitempty"`

	Source *Source `json:"source"`

	ResourceLimit *ResourceLimit `json:"resource_limit"`

	// 镜像地址。
	ImageUrl *string `json:"image_url,omitempty"`
}

func (o CreateComponentWithConfigurationRequestBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentWithConfigurationRequestBodySpec struct{}"
	}

	return strings.Join([]string{"CreateComponentWithConfigurationRequestBodySpec", string(data)}, " ")
}

type CreateComponentWithConfigurationRequestBodySpecRuntime struct {
	value string
}

type CreateComponentWithConfigurationRequestBodySpecRuntimeEnum struct {
	DOCKER   CreateComponentWithConfigurationRequestBodySpecRuntime
	JAVA8    CreateComponentWithConfigurationRequestBodySpecRuntime
	JAVA11   CreateComponentWithConfigurationRequestBodySpecRuntime
	JAVA17   CreateComponentWithConfigurationRequestBodySpecRuntime
	TOMCAT8  CreateComponentWithConfigurationRequestBodySpecRuntime
	TOMCAT9  CreateComponentWithConfigurationRequestBodySpecRuntime
	PYTHON3  CreateComponentWithConfigurationRequestBodySpecRuntime
	NODEJS8  CreateComponentWithConfigurationRequestBodySpecRuntime
	NODEJS14 CreateComponentWithConfigurationRequestBodySpecRuntime
	NODEJS16 CreateComponentWithConfigurationRequestBodySpecRuntime
	PHP7     CreateComponentWithConfigurationRequestBodySpecRuntime
	PHP8     CreateComponentWithConfigurationRequestBodySpecRuntime
	DOTNET6  CreateComponentWithConfigurationRequestBodySpecRuntime
	DOTNET7  CreateComponentWithConfigurationRequestBodySpecRuntime
	DOTNET8  CreateComponentWithConfigurationRequestBodySpecRuntime
}

func GetCreateComponentWithConfigurationRequestBodySpecRuntimeEnum() CreateComponentWithConfigurationRequestBodySpecRuntimeEnum {
	return CreateComponentWithConfigurationRequestBodySpecRuntimeEnum{
		DOCKER: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Docker",
		},
		JAVA8: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Java8",
		},
		JAVA11: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Java11",
		},
		JAVA17: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Java17",
		},
		TOMCAT8: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Tomcat8",
		},
		TOMCAT9: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Tomcat9",
		},
		PYTHON3: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Python3",
		},
		NODEJS8: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Nodejs8",
		},
		NODEJS14: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Nodejs14",
		},
		NODEJS16: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Nodejs16",
		},
		PHP7: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Php7",
		},
		PHP8: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Php8",
		},
		DOTNET6: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Dotnet6",
		},
		DOTNET7: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Dotnet7",
		},
		DOTNET8: CreateComponentWithConfigurationRequestBodySpecRuntime{
			value: "Dotnet8",
		},
	}
}

func (c CreateComponentWithConfigurationRequestBodySpecRuntime) Value() string {
	return c.value
}

func (c CreateComponentWithConfigurationRequestBodySpecRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentWithConfigurationRequestBodySpecRuntime) UnmarshalJSON(b []byte) error {
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
