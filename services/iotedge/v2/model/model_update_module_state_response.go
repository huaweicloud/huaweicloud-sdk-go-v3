package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateModuleStateResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 模块运行状态
	State *UpdateModuleStateResponseState `json:"state,omitempty"`

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
	AppType *UpdateModuleStateResponseAppType `json:"app_type,omitempty"`

	// 功能类型
	FunctionType   *UpdateModuleStateResponseFunctionType `json:"function_type,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o UpdateModuleStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateModuleStateResponse", string(data)}, " ")
}

type UpdateModuleStateResponseState struct {
	value string
}

type UpdateModuleStateResponseStateEnum struct {
	PENDING        UpdateModuleStateResponseState
	PENDING_DELETE UpdateModuleStateResponseState
	DELETE_FAILED  UpdateModuleStateResponseState
	RUNNING        UpdateModuleStateResponseState
	FAILED         UpdateModuleStateResponseState
	SUCCEEDED      UpdateModuleStateResponseState
	UNKNOWN        UpdateModuleStateResponseState
	DELETE_SUCCESS UpdateModuleStateResponseState
	STOPPED        UpdateModuleStateResponseState
}

func GetUpdateModuleStateResponseStateEnum() UpdateModuleStateResponseStateEnum {
	return UpdateModuleStateResponseStateEnum{
		PENDING: UpdateModuleStateResponseState{
			value: "PENDING",
		},
		PENDING_DELETE: UpdateModuleStateResponseState{
			value: "PENDING_DELETE",
		},
		DELETE_FAILED: UpdateModuleStateResponseState{
			value: "DELETE_FAILED",
		},
		RUNNING: UpdateModuleStateResponseState{
			value: "RUNNING",
		},
		FAILED: UpdateModuleStateResponseState{
			value: "FAILED",
		},
		SUCCEEDED: UpdateModuleStateResponseState{
			value: "SUCCEEDED",
		},
		UNKNOWN: UpdateModuleStateResponseState{
			value: "UNKNOWN",
		},
		DELETE_SUCCESS: UpdateModuleStateResponseState{
			value: "DELETE_SUCCESS",
		},
		STOPPED: UpdateModuleStateResponseState{
			value: "STOPPED",
		},
	}
}

func (c UpdateModuleStateResponseState) Value() string {
	return c.value
}

func (c UpdateModuleStateResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateModuleStateResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateModuleStateResponseAppType struct {
	value string
}

type UpdateModuleStateResponseAppTypeEnum struct {
	SYSTEM_REQUIRED UpdateModuleStateResponseAppType
	SYSTEM_OPTIONAL UpdateModuleStateResponseAppType
	USER            UpdateModuleStateResponseAppType
}

func GetUpdateModuleStateResponseAppTypeEnum() UpdateModuleStateResponseAppTypeEnum {
	return UpdateModuleStateResponseAppTypeEnum{
		SYSTEM_REQUIRED: UpdateModuleStateResponseAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: UpdateModuleStateResponseAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: UpdateModuleStateResponseAppType{
			value: "USER",
		},
	}
}

func (c UpdateModuleStateResponseAppType) Value() string {
	return c.value
}

func (c UpdateModuleStateResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateModuleStateResponseAppType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateModuleStateResponseFunctionType struct {
	value string
}

type UpdateModuleStateResponseFunctionTypeEnum struct {
	DATA_PROCESSING        UpdateModuleStateResponseFunctionType
	PROTOCOL_PARSING       UpdateModuleStateResponseFunctionType
	ON_PREMISE_INTEGRATION UpdateModuleStateResponseFunctionType
	GATEWAY_MANAGER        UpdateModuleStateResponseFunctionType
	COMPOSITE_APPLICATION  UpdateModuleStateResponseFunctionType
	DATA_COLLECTION        UpdateModuleStateResponseFunctionType
}

func GetUpdateModuleStateResponseFunctionTypeEnum() UpdateModuleStateResponseFunctionTypeEnum {
	return UpdateModuleStateResponseFunctionTypeEnum{
		DATA_PROCESSING: UpdateModuleStateResponseFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: UpdateModuleStateResponseFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: UpdateModuleStateResponseFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
		GATEWAY_MANAGER: UpdateModuleStateResponseFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: UpdateModuleStateResponseFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: UpdateModuleStateResponseFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c UpdateModuleStateResponseFunctionType) Value() string {
	return c.value
}

func (c UpdateModuleStateResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateModuleStateResponseFunctionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
