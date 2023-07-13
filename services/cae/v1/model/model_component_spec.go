package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentSpec 组件规格。
type ComponentSpec struct {

	// 语言/运行时。
	Runtime *ComponentSpecRuntime `json:"runtime,omitempty"`

	// 环境ID。
	EnvId *string `json:"env_id,omitempty"`

	// 实例个数。
	Replica *int32 `json:"replica,omitempty"`

	Source *Source `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`

	ResourceLimit *ResourceLimit `json:"resource_limit,omitempty"`

	// 访问方式列表。
	AccessInfo *[]Access `json:"access_info,omitempty"`

	// 镜像地址。
	ImageUrl *string `json:"image_url,omitempty"`

	// 可用实例个数。
	AvailableReplica *int32 `json:"available_replica,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 构建任务ID。
	BuildId *string `json:"build_id,omitempty"`

	// 组件状态。
	Status *string `json:"status,omitempty"`

	// 构建日志ID。
	BuildLogId *string `json:"build_log_id,omitempty"`
}

func (o ComponentSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentSpec struct{}"
	}

	return strings.Join([]string{"ComponentSpec", string(data)}, " ")
}

type ComponentSpecRuntime struct {
	value string
}

type ComponentSpecRuntimeEnum struct {
	DOCKER  ComponentSpecRuntime
	JAVA8   ComponentSpecRuntime
	JAVA11  ComponentSpecRuntime
	TOMCAT8 ComponentSpecRuntime
	TOMCAT9 ComponentSpecRuntime
	PYTHON3 ComponentSpecRuntime
	NODEJS8 ComponentSpecRuntime
	PHP7    ComponentSpecRuntime
}

func GetComponentSpecRuntimeEnum() ComponentSpecRuntimeEnum {
	return ComponentSpecRuntimeEnum{
		DOCKER: ComponentSpecRuntime{
			value: "Docker",
		},
		JAVA8: ComponentSpecRuntime{
			value: "Java8",
		},
		JAVA11: ComponentSpecRuntime{
			value: "Java11",
		},
		TOMCAT8: ComponentSpecRuntime{
			value: "Tomcat8",
		},
		TOMCAT9: ComponentSpecRuntime{
			value: "Tomcat9",
		},
		PYTHON3: ComponentSpecRuntime{
			value: "Python3",
		},
		NODEJS8: ComponentSpecRuntime{
			value: "Nodejs8",
		},
		PHP7: ComponentSpecRuntime{
			value: "Php7",
		},
	}
}

func (c ComponentSpecRuntime) Value() string {
	return c.value
}

func (c ComponentSpecRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentSpecRuntime) UnmarshalJSON(b []byte) error {
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
