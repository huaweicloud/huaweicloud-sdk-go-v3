package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateComponentSpec 组件规格。
type CreateComponentSpec struct {

	// 语言/运行时。
	Runtime *CreateComponentSpecRuntime `json:"runtime,omitempty"`

	// 环境ID。
	EnvId *string `json:"env_id,omitempty"`

	// 实例个数。
	Replica *int32 `json:"replica,omitempty"`

	Source *Source `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`

	ResourceLimit *ResourceLimit `json:"resource_limit,omitempty"`

	// 可用实例个数。
	AvailableReplica *int32 `json:"available_replica,omitempty"`

	// 组件状态。
	Status *string `json:"status,omitempty"`
}

func (o CreateComponentSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentSpec struct{}"
	}

	return strings.Join([]string{"CreateComponentSpec", string(data)}, " ")
}

type CreateComponentSpecRuntime struct {
	value string
}

type CreateComponentSpecRuntimeEnum struct {
	DOCKER  CreateComponentSpecRuntime
	JAVA8   CreateComponentSpecRuntime
	JAVA11  CreateComponentSpecRuntime
	TOMCAT8 CreateComponentSpecRuntime
	TOMCAT9 CreateComponentSpecRuntime
	PYTHON3 CreateComponentSpecRuntime
	NODEJS8 CreateComponentSpecRuntime
	PHP7    CreateComponentSpecRuntime
}

func GetCreateComponentSpecRuntimeEnum() CreateComponentSpecRuntimeEnum {
	return CreateComponentSpecRuntimeEnum{
		DOCKER: CreateComponentSpecRuntime{
			value: "Docker",
		},
		JAVA8: CreateComponentSpecRuntime{
			value: "Java8",
		},
		JAVA11: CreateComponentSpecRuntime{
			value: "Java11",
		},
		TOMCAT8: CreateComponentSpecRuntime{
			value: "Tomcat8",
		},
		TOMCAT9: CreateComponentSpecRuntime{
			value: "Tomcat9",
		},
		PYTHON3: CreateComponentSpecRuntime{
			value: "Python3",
		},
		NODEJS8: CreateComponentSpecRuntime{
			value: "Nodejs8",
		},
		PHP7: CreateComponentSpecRuntime{
			value: "Php7",
		},
	}
}

func (c CreateComponentSpecRuntime) Value() string {
	return c.value
}

func (c CreateComponentSpecRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentSpecRuntime) UnmarshalJSON(b []byte) error {
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
