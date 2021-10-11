package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowModuleResponse struct {
	// 应用ID

	EdgeAppId *string `json:"edge_app_id,omitempty"`
	// 应用版本

	AppVersion *string `json:"app_version,omitempty"`
	// 模块状态

	State *ShowModuleResponseState `json:"state,omitempty"`
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

	AppType *ShowModuleResponseAppType `json:"app_type,omitempty"`
	// 功能类型

	FunctionType   *ShowModuleResponseFunctionType `json:"function_type,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowModuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowModuleResponse struct{}"
	}

	return strings.Join([]string{"ShowModuleResponse", string(data)}, " ")
}

type ShowModuleResponseState struct {
	value string
}

type ShowModuleResponseStateEnum struct {
	PENDING        ShowModuleResponseState
	PENDING_DELETE ShowModuleResponseState
	DELETE_FAILED  ShowModuleResponseState
	RUNNING        ShowModuleResponseState
	FAILED         ShowModuleResponseState
	SUCCEEDED      ShowModuleResponseState
	UNKNOWN        ShowModuleResponseState
}

func GetShowModuleResponseStateEnum() ShowModuleResponseStateEnum {
	return ShowModuleResponseStateEnum{
		PENDING: ShowModuleResponseState{
			value: "PENDING",
		},
		PENDING_DELETE: ShowModuleResponseState{
			value: "PENDING_DELETE",
		},
		DELETE_FAILED: ShowModuleResponseState{
			value: "DELETE_FAILED",
		},
		RUNNING: ShowModuleResponseState{
			value: "RUNNING",
		},
		FAILED: ShowModuleResponseState{
			value: "FAILED",
		},
		SUCCEEDED: ShowModuleResponseState{
			value: "SUCCEEDED",
		},
		UNKNOWN: ShowModuleResponseState{
			value: "UNKNOWN",
		},
	}
}

func (c ShowModuleResponseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowModuleResponseState) UnmarshalJSON(b []byte) error {
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

type ShowModuleResponseAppType struct {
	value string
}

type ShowModuleResponseAppTypeEnum struct {
	SYSTEM_REQUIRED ShowModuleResponseAppType
	SYSTEM_OPTIONAL ShowModuleResponseAppType
	USER            ShowModuleResponseAppType
}

func GetShowModuleResponseAppTypeEnum() ShowModuleResponseAppTypeEnum {
	return ShowModuleResponseAppTypeEnum{
		SYSTEM_REQUIRED: ShowModuleResponseAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: ShowModuleResponseAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: ShowModuleResponseAppType{
			value: "USER",
		},
	}
}

func (c ShowModuleResponseAppType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowModuleResponseAppType) UnmarshalJSON(b []byte) error {
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

type ShowModuleResponseFunctionType struct {
	value string
}

type ShowModuleResponseFunctionTypeEnum struct {
	DATA_PROCESSING        ShowModuleResponseFunctionType
	PROTOCOL_PARSING       ShowModuleResponseFunctionType
	ON_PREMISE_INTEGRATION ShowModuleResponseFunctionType
}

func GetShowModuleResponseFunctionTypeEnum() ShowModuleResponseFunctionTypeEnum {
	return ShowModuleResponseFunctionTypeEnum{
		DATA_PROCESSING: ShowModuleResponseFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: ShowModuleResponseFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: ShowModuleResponseFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
	}
}

func (c ShowModuleResponseFunctionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowModuleResponseFunctionType) UnmarshalJSON(b []byte) error {
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
