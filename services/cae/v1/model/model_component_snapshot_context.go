package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComponentSnapshotContext struct {

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 可用实例个数。
	AvailableReplica *int32 `json:"available_replica,omitempty"`

	// 组件构建信息。
	Build *string `json:"build,omitempty"`

	// 构建任务ID。
	BuildId *string `json:"build_id,omitempty"`

	// 构建日志ID。
	BuildLogId *string `json:"build_log_id,omitempty"`

	// 环境ID。
	EnvId *string `json:"env_id,omitempty"`

	// 组件ID。
	Id *string `json:"id,omitempty"`

	// 镜像地址。
	ImageUrl *string `json:"image_url,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// LTS日志组的ID。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// LTS日志流的ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 组件名称。
	Name *string `json:"name,omitempty"`

	// 组件操作。
	Operation *string `json:"operation,omitempty"`

	// 组件操作状态。
	OperationStatus *string `json:"operation_status,omitempty"`

	// 实例个数。
	Replica *int32 `json:"replica,omitempty"`

	// 组件规格。
	ResourceLimit *string `json:"resource_limit,omitempty"`

	// 语言/运行时。
	Runtime *ComponentSnapshotContextRuntime `json:"runtime,omitempty"`

	// 组件源信息。
	Source *string `json:"source,omitempty"`

	// 组件状态。
	Status *string `json:"status,omitempty"`

	// 组件版本。
	Version *string `json:"version,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ComponentSnapshotContext) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentSnapshotContext struct{}"
	}

	return strings.Join([]string{"ComponentSnapshotContext", string(data)}, " ")
}

type ComponentSnapshotContextRuntime struct {
	value string
}

type ComponentSnapshotContextRuntimeEnum struct {
	DOCKER   ComponentSnapshotContextRuntime
	JAVA8    ComponentSnapshotContextRuntime
	JAVA11   ComponentSnapshotContextRuntime
	JAVA17   ComponentSnapshotContextRuntime
	TOMCAT8  ComponentSnapshotContextRuntime
	TOMCAT9  ComponentSnapshotContextRuntime
	PYTHON3  ComponentSnapshotContextRuntime
	NODEJS8  ComponentSnapshotContextRuntime
	NODEJS14 ComponentSnapshotContextRuntime
	NODEJS16 ComponentSnapshotContextRuntime
	PHP7     ComponentSnapshotContextRuntime
	PHP8     ComponentSnapshotContextRuntime
	DOTNET6  ComponentSnapshotContextRuntime
	DOTNET7  ComponentSnapshotContextRuntime
	DOTNET8  ComponentSnapshotContextRuntime
}

func GetComponentSnapshotContextRuntimeEnum() ComponentSnapshotContextRuntimeEnum {
	return ComponentSnapshotContextRuntimeEnum{
		DOCKER: ComponentSnapshotContextRuntime{
			value: "Docker",
		},
		JAVA8: ComponentSnapshotContextRuntime{
			value: "Java8",
		},
		JAVA11: ComponentSnapshotContextRuntime{
			value: "Java11",
		},
		JAVA17: ComponentSnapshotContextRuntime{
			value: "Java17",
		},
		TOMCAT8: ComponentSnapshotContextRuntime{
			value: "Tomcat8",
		},
		TOMCAT9: ComponentSnapshotContextRuntime{
			value: "Tomcat9",
		},
		PYTHON3: ComponentSnapshotContextRuntime{
			value: "Python3",
		},
		NODEJS8: ComponentSnapshotContextRuntime{
			value: "Nodejs8",
		},
		NODEJS14: ComponentSnapshotContextRuntime{
			value: "Nodejs14",
		},
		NODEJS16: ComponentSnapshotContextRuntime{
			value: "Nodejs16",
		},
		PHP7: ComponentSnapshotContextRuntime{
			value: "Php7",
		},
		PHP8: ComponentSnapshotContextRuntime{
			value: "Php8",
		},
		DOTNET6: ComponentSnapshotContextRuntime{
			value: "Dotnet6",
		},
		DOTNET7: ComponentSnapshotContextRuntime{
			value: "Dotnet7",
		},
		DOTNET8: ComponentSnapshotContextRuntime{
			value: "Dotnet8",
		},
	}
}

func (c ComponentSnapshotContextRuntime) Value() string {
	return c.value
}

func (c ComponentSnapshotContextRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentSnapshotContextRuntime) UnmarshalJSON(b []byte) error {
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
