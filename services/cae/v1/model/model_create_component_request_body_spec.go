package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateComponentRequestBodySpec 组件规格信息。
type CreateComponentRequestBodySpec struct {

	// 语言/运行时。
	Runtime CreateComponentRequestBodySpecRuntime `json:"runtime"`

	// 实例个数。
	Replica int32 `json:"replica"`

	Build *Build `json:"build,omitempty"`

	Source *Source `json:"source"`

	ResourceLimit *ResourceLimit `json:"resource_limit"`

	// 镜像地址。
	ImageUrl *string `json:"image_url,omitempty"`
}

func (o CreateComponentRequestBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequestBodySpec struct{}"
	}

	return strings.Join([]string{"CreateComponentRequestBodySpec", string(data)}, " ")
}

type CreateComponentRequestBodySpecRuntime struct {
	value string
}

type CreateComponentRequestBodySpecRuntimeEnum struct {
	DOCKER  CreateComponentRequestBodySpecRuntime
	JAVA8   CreateComponentRequestBodySpecRuntime
	JAVA11  CreateComponentRequestBodySpecRuntime
	TOMCAT8 CreateComponentRequestBodySpecRuntime
	TOMCAT9 CreateComponentRequestBodySpecRuntime
	PYTHON3 CreateComponentRequestBodySpecRuntime
	NODEJS8 CreateComponentRequestBodySpecRuntime
	PHP7    CreateComponentRequestBodySpecRuntime
}

func GetCreateComponentRequestBodySpecRuntimeEnum() CreateComponentRequestBodySpecRuntimeEnum {
	return CreateComponentRequestBodySpecRuntimeEnum{
		DOCKER: CreateComponentRequestBodySpecRuntime{
			value: "Docker",
		},
		JAVA8: CreateComponentRequestBodySpecRuntime{
			value: "Java8",
		},
		JAVA11: CreateComponentRequestBodySpecRuntime{
			value: "Java11",
		},
		TOMCAT8: CreateComponentRequestBodySpecRuntime{
			value: "Tomcat8",
		},
		TOMCAT9: CreateComponentRequestBodySpecRuntime{
			value: "Tomcat9",
		},
		PYTHON3: CreateComponentRequestBodySpecRuntime{
			value: "Python3",
		},
		NODEJS8: CreateComponentRequestBodySpecRuntime{
			value: "Nodejs8",
		},
		PHP7: CreateComponentRequestBodySpecRuntime{
			value: "Php7",
		},
	}
}

func (c CreateComponentRequestBodySpecRuntime) Value() string {
	return c.value
}

func (c CreateComponentRequestBodySpecRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentRequestBodySpecRuntime) UnmarshalJSON(b []byte) error {
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
