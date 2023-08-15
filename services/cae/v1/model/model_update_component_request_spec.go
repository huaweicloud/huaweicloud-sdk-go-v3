package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateComponentRequestSpec struct {

	// 语言/运行时。
	Runtime *UpdateComponentRequestSpecRuntime `json:"runtime,omitempty"`

	Source *Source `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`

	ResourceLimit *ResourceLimit `json:"resource_limit"`

	// 实例个数。
	Replica *int32 `json:"replica,omitempty"`
}

func (o UpdateComponentRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequestSpec struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequestSpec", string(data)}, " ")
}

type UpdateComponentRequestSpecRuntime struct {
	value string
}

type UpdateComponentRequestSpecRuntimeEnum struct {
	DOCKER  UpdateComponentRequestSpecRuntime
	JAVA8   UpdateComponentRequestSpecRuntime
	JAVA11  UpdateComponentRequestSpecRuntime
	TOMCAT8 UpdateComponentRequestSpecRuntime
	TOMCAT9 UpdateComponentRequestSpecRuntime
	PYTHON3 UpdateComponentRequestSpecRuntime
	NODEJS8 UpdateComponentRequestSpecRuntime
	PHP7    UpdateComponentRequestSpecRuntime
}

func GetUpdateComponentRequestSpecRuntimeEnum() UpdateComponentRequestSpecRuntimeEnum {
	return UpdateComponentRequestSpecRuntimeEnum{
		DOCKER: UpdateComponentRequestSpecRuntime{
			value: "Docker",
		},
		JAVA8: UpdateComponentRequestSpecRuntime{
			value: "Java8",
		},
		JAVA11: UpdateComponentRequestSpecRuntime{
			value: "Java11",
		},
		TOMCAT8: UpdateComponentRequestSpecRuntime{
			value: "Tomcat8",
		},
		TOMCAT9: UpdateComponentRequestSpecRuntime{
			value: "Tomcat9",
		},
		PYTHON3: UpdateComponentRequestSpecRuntime{
			value: "Python3",
		},
		NODEJS8: UpdateComponentRequestSpecRuntime{
			value: "Nodejs8",
		},
		PHP7: UpdateComponentRequestSpecRuntime{
			value: "Php7",
		},
	}
}

func (c UpdateComponentRequestSpecRuntime) Value() string {
	return c.value
}

func (c UpdateComponentRequestSpecRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateComponentRequestSpecRuntime) UnmarshalJSON(b []byte) error {
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
