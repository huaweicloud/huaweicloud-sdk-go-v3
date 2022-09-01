package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateModuleResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty" xml:"edge_app_id"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty" xml:"app_version"`

	// 模块运行状态
	State *UpdateModuleResponseState `json:"state,omitempty" xml:"state"`

	// 模块管控状态
	ControlStatus *string `json:"control_status,omitempty" xml:"control_status"`

	// 边缘节点（同deviceID）ID
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty" xml:"module_name"`

	// 模块ID
	ModuleId *string `json:"module_id,omitempty" xml:"module_id"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 应用类型
	AppType *UpdateModuleResponseAppType `json:"app_type,omitempty" xml:"app_type"`

	// 功能类型
	FunctionType   *UpdateModuleResponseFunctionType `json:"function_type,omitempty" xml:"function_type"`
	HttpStatusCode int                               `json:"-"`
}

func (o UpdateModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateModuleResponse", string(data)}, " ")
}

type UpdateModuleResponseState struct {
	value string
}

type UpdateModuleResponseStateEnum struct {
	PENDING        UpdateModuleResponseState
	PENDING_DELETE UpdateModuleResponseState
	DELETE_FAILED  UpdateModuleResponseState
	RUNNING        UpdateModuleResponseState
	FAILED         UpdateModuleResponseState
	SUCCEEDED      UpdateModuleResponseState
	UNKNOWN        UpdateModuleResponseState
	DELETE_SUCCESS UpdateModuleResponseState
	STOPPED        UpdateModuleResponseState
}

func GetUpdateModuleResponseStateEnum() UpdateModuleResponseStateEnum {
	return UpdateModuleResponseStateEnum{
		PENDING: UpdateModuleResponseState{
			value: "PENDING",
		},
		PENDING_DELETE: UpdateModuleResponseState{
			value: "PENDING_DELETE",
		},
		DELETE_FAILED: UpdateModuleResponseState{
			value: "DELETE_FAILED",
		},
		RUNNING: UpdateModuleResponseState{
			value: "RUNNING",
		},
		FAILED: UpdateModuleResponseState{
			value: "FAILED",
		},
		SUCCEEDED: UpdateModuleResponseState{
			value: "SUCCEEDED",
		},
		UNKNOWN: UpdateModuleResponseState{
			value: "UNKNOWN",
		},
		DELETE_SUCCESS: UpdateModuleResponseState{
			value: "DELETE_SUCCESS",
		},
		STOPPED: UpdateModuleResponseState{
			value: "STOPPED",
		},
	}
}

func (c UpdateModuleResponseState) Value() string {
	return c.value
}

func (c UpdateModuleResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateModuleResponseState) UnmarshalJSON(b []byte) error {
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

type UpdateModuleResponseAppType struct {
	value string
}

type UpdateModuleResponseAppTypeEnum struct {
	SYSTEM_REQUIRED UpdateModuleResponseAppType
	SYSTEM_OPTIONAL UpdateModuleResponseAppType
	USER            UpdateModuleResponseAppType
}

func GetUpdateModuleResponseAppTypeEnum() UpdateModuleResponseAppTypeEnum {
	return UpdateModuleResponseAppTypeEnum{
		SYSTEM_REQUIRED: UpdateModuleResponseAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: UpdateModuleResponseAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: UpdateModuleResponseAppType{
			value: "USER",
		},
	}
}

func (c UpdateModuleResponseAppType) Value() string {
	return c.value
}

func (c UpdateModuleResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateModuleResponseAppType) UnmarshalJSON(b []byte) error {
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

type UpdateModuleResponseFunctionType struct {
	value string
}

type UpdateModuleResponseFunctionTypeEnum struct {
	DATA_PROCESSING        UpdateModuleResponseFunctionType
	PROTOCOL_PARSING       UpdateModuleResponseFunctionType
	ON_PREMISE_INTEGRATION UpdateModuleResponseFunctionType
	GATEWAY_MANAGER        UpdateModuleResponseFunctionType
	COMPOSITE_APPLICATION  UpdateModuleResponseFunctionType
	DATA_COLLECTION        UpdateModuleResponseFunctionType
}

func GetUpdateModuleResponseFunctionTypeEnum() UpdateModuleResponseFunctionTypeEnum {
	return UpdateModuleResponseFunctionTypeEnum{
		DATA_PROCESSING: UpdateModuleResponseFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: UpdateModuleResponseFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: UpdateModuleResponseFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
		GATEWAY_MANAGER: UpdateModuleResponseFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: UpdateModuleResponseFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: UpdateModuleResponseFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c UpdateModuleResponseFunctionType) Value() string {
	return c.value
}

func (c UpdateModuleResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateModuleResponseFunctionType) UnmarshalJSON(b []byte) error {
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
