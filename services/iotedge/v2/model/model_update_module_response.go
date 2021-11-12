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

	EdgeAppId *string `json:"edge_app_id,omitempty"`
	// 应用版本

	AppVersion *string `json:"app_version,omitempty"`
	// 模块状态

	State *UpdateModuleResponseState `json:"state,omitempty"`
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

	AppType *UpdateModuleResponseAppType `json:"app_type,omitempty"`
	// 功能类型

	FunctionType   *UpdateModuleResponseFunctionType `json:"function_type,omitempty"`
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
	}
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
	}
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
