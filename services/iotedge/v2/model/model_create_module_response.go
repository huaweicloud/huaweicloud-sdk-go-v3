package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateModuleResponse Response Object
type CreateModuleResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 模块运行状态
	State *CreateModuleResponseState `json:"state,omitempty"`

	// 模块管控状态
	ControlStatus *string `json:"control_status,omitempty"`

	// 边缘节点（同deviceID）ID
	NodeId *string `json:"node_id,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`

	// 模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 应用类型
	AppType *CreateModuleResponseAppType `json:"app_type,omitempty"`

	// 功能类型
	FunctionType *CreateModuleResponseFunctionType `json:"function_type,omitempty"`

	ContainerSettings *ModuleContainerSettingsResDto `json:"container_settings,omitempty"`
	HttpStatusCode    int                            `json:"-"`
}

func (o CreateModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModuleResponse struct{}"
	}

	return strings.Join([]string{"CreateModuleResponse", string(data)}, " ")
}

type CreateModuleResponseState struct {
	value string
}

type CreateModuleResponseStateEnum struct {
	PENDING        CreateModuleResponseState
	PENDING_DELETE CreateModuleResponseState
	DELETE_FAILED  CreateModuleResponseState
	RUNNING        CreateModuleResponseState
	FAILED         CreateModuleResponseState
	SUCCEEDED      CreateModuleResponseState
	UNKNOWN        CreateModuleResponseState
	DELETE_SUCCESS CreateModuleResponseState
	STOPPED        CreateModuleResponseState
}

func GetCreateModuleResponseStateEnum() CreateModuleResponseStateEnum {
	return CreateModuleResponseStateEnum{
		PENDING: CreateModuleResponseState{
			value: "PENDING",
		},
		PENDING_DELETE: CreateModuleResponseState{
			value: "PENDING_DELETE",
		},
		DELETE_FAILED: CreateModuleResponseState{
			value: "DELETE_FAILED",
		},
		RUNNING: CreateModuleResponseState{
			value: "RUNNING",
		},
		FAILED: CreateModuleResponseState{
			value: "FAILED",
		},
		SUCCEEDED: CreateModuleResponseState{
			value: "SUCCEEDED",
		},
		UNKNOWN: CreateModuleResponseState{
			value: "UNKNOWN",
		},
		DELETE_SUCCESS: CreateModuleResponseState{
			value: "DELETE_SUCCESS",
		},
		STOPPED: CreateModuleResponseState{
			value: "STOPPED",
		},
	}
}

func (c CreateModuleResponseState) Value() string {
	return c.value
}

func (c CreateModuleResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateModuleResponseState) UnmarshalJSON(b []byte) error {
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

type CreateModuleResponseAppType struct {
	value string
}

type CreateModuleResponseAppTypeEnum struct {
	SYSTEM_REQUIRED CreateModuleResponseAppType
	SYSTEM_OPTIONAL CreateModuleResponseAppType
	USER            CreateModuleResponseAppType
}

func GetCreateModuleResponseAppTypeEnum() CreateModuleResponseAppTypeEnum {
	return CreateModuleResponseAppTypeEnum{
		SYSTEM_REQUIRED: CreateModuleResponseAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: CreateModuleResponseAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: CreateModuleResponseAppType{
			value: "USER",
		},
	}
}

func (c CreateModuleResponseAppType) Value() string {
	return c.value
}

func (c CreateModuleResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateModuleResponseAppType) UnmarshalJSON(b []byte) error {
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

type CreateModuleResponseFunctionType struct {
	value string
}

type CreateModuleResponseFunctionTypeEnum struct {
	DATA_PROCESSING        CreateModuleResponseFunctionType
	PROTOCOL_PARSING       CreateModuleResponseFunctionType
	ON_PREMISE_INTEGRATION CreateModuleResponseFunctionType
	GATEWAY_MANAGER        CreateModuleResponseFunctionType
	COMPOSITE_APPLICATION  CreateModuleResponseFunctionType
	DATA_COLLECTION        CreateModuleResponseFunctionType
}

func GetCreateModuleResponseFunctionTypeEnum() CreateModuleResponseFunctionTypeEnum {
	return CreateModuleResponseFunctionTypeEnum{
		DATA_PROCESSING: CreateModuleResponseFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: CreateModuleResponseFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: CreateModuleResponseFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
		GATEWAY_MANAGER: CreateModuleResponseFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: CreateModuleResponseFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: CreateModuleResponseFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c CreateModuleResponseFunctionType) Value() string {
	return c.value
}

func (c CreateModuleResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateModuleResponseFunctionType) UnmarshalJSON(b []byte) error {
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
