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
	Status *ComponentSpecStatus `json:"status,omitempty"`

	// 构建日志ID。
	BuildLogId *string `json:"build_log_id,omitempty"`

	// 组件最新配置的操作ID。
	ConfigurationOperationId *string `json:"configuration_operation_id,omitempty"`
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
	DOCKER   ComponentSpecRuntime
	JAVA8    ComponentSpecRuntime
	JAVA11   ComponentSpecRuntime
	JAVA17   ComponentSpecRuntime
	TOMCAT8  ComponentSpecRuntime
	TOMCAT9  ComponentSpecRuntime
	PYTHON3  ComponentSpecRuntime
	NODEJS8  ComponentSpecRuntime
	NODEJS14 ComponentSpecRuntime
	NODEJS16 ComponentSpecRuntime
	PHP7     ComponentSpecRuntime
	PHP8     ComponentSpecRuntime
	DOTNET6  ComponentSpecRuntime
	DOTNET7  ComponentSpecRuntime
	DOTNET8  ComponentSpecRuntime
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
		JAVA17: ComponentSpecRuntime{
			value: "Java17",
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
		NODEJS14: ComponentSpecRuntime{
			value: "Nodejs14",
		},
		NODEJS16: ComponentSpecRuntime{
			value: "Nodejs16",
		},
		PHP7: ComponentSpecRuntime{
			value: "Php7",
		},
		PHP8: ComponentSpecRuntime{
			value: "Php8",
		},
		DOTNET6: ComponentSpecRuntime{
			value: "Dotnet6",
		},
		DOTNET7: ComponentSpecRuntime{
			value: "Dotnet7",
		},
		DOTNET8: ComponentSpecRuntime{
			value: "Dotnet8",
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

type ComponentSpecStatus struct {
	value string
}

type ComponentSpecStatusEnum struct {
	CREATED   ComponentSpecStatus
	RUNNING   ComponentSpecStatus
	PAUSED    ComponentSpecStatus
	NOT_READY ComponentSpecStatus
}

func GetComponentSpecStatusEnum() ComponentSpecStatusEnum {
	return ComponentSpecStatusEnum{
		CREATED: ComponentSpecStatus{
			value: "created",
		},
		RUNNING: ComponentSpecStatus{
			value: "running",
		},
		PAUSED: ComponentSpecStatus{
			value: "paused",
		},
		NOT_READY: ComponentSpecStatus{
			value: "notReady",
		},
	}
}

func (c ComponentSpecStatus) Value() string {
	return c.value
}

func (c ComponentSpecStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentSpecStatus) UnmarshalJSON(b []byte) error {
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
