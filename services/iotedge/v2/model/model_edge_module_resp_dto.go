package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建边缘模块响应结构体
type EdgeModuleRespDto struct {
	// 应用ID

	EdgeAppId *string `json:"edge_app_id,omitempty"`
	// 应用版本

	AppVersion *string `json:"app_version,omitempty"`
	// 模块状态

	State *EdgeModuleRespDtoState `json:"state,omitempty"`
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

	AppType *EdgeModuleRespDtoAppType `json:"app_type,omitempty"`
	// 功能类型

	FunctionType *EdgeModuleRespDtoFunctionType `json:"function_type,omitempty"`
}

func (o EdgeModuleRespDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EdgeModuleRespDto struct{}"
	}

	return strings.Join([]string{"EdgeModuleRespDto", string(data)}, " ")
}

type EdgeModuleRespDtoState struct {
	value string
}

type EdgeModuleRespDtoStateEnum struct {
	PENDING        EdgeModuleRespDtoState
	PENDING_DELETE EdgeModuleRespDtoState
	DELETE_FAILED  EdgeModuleRespDtoState
	RUNNING        EdgeModuleRespDtoState
	FAILED         EdgeModuleRespDtoState
	SUCCEEDED      EdgeModuleRespDtoState
	UNKNOWN        EdgeModuleRespDtoState
}

func GetEdgeModuleRespDtoStateEnum() EdgeModuleRespDtoStateEnum {
	return EdgeModuleRespDtoStateEnum{
		PENDING: EdgeModuleRespDtoState{
			value: "PENDING",
		},
		PENDING_DELETE: EdgeModuleRespDtoState{
			value: "PENDING_DELETE",
		},
		DELETE_FAILED: EdgeModuleRespDtoState{
			value: "DELETE_FAILED",
		},
		RUNNING: EdgeModuleRespDtoState{
			value: "RUNNING",
		},
		FAILED: EdgeModuleRespDtoState{
			value: "FAILED",
		},
		SUCCEEDED: EdgeModuleRespDtoState{
			value: "SUCCEEDED",
		},
		UNKNOWN: EdgeModuleRespDtoState{
			value: "UNKNOWN",
		},
	}
}

func (c EdgeModuleRespDtoState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *EdgeModuleRespDtoState) UnmarshalJSON(b []byte) error {
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

type EdgeModuleRespDtoAppType struct {
	value string
}

type EdgeModuleRespDtoAppTypeEnum struct {
	SYSTEM_REQUIRED EdgeModuleRespDtoAppType
	SYSTEM_OPTIONAL EdgeModuleRespDtoAppType
	USER            EdgeModuleRespDtoAppType
}

func GetEdgeModuleRespDtoAppTypeEnum() EdgeModuleRespDtoAppTypeEnum {
	return EdgeModuleRespDtoAppTypeEnum{
		SYSTEM_REQUIRED: EdgeModuleRespDtoAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: EdgeModuleRespDtoAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: EdgeModuleRespDtoAppType{
			value: "USER",
		},
	}
}

func (c EdgeModuleRespDtoAppType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *EdgeModuleRespDtoAppType) UnmarshalJSON(b []byte) error {
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

type EdgeModuleRespDtoFunctionType struct {
	value string
}

type EdgeModuleRespDtoFunctionTypeEnum struct {
	DATA_PROCESSING        EdgeModuleRespDtoFunctionType
	PROTOCOL_PARSING       EdgeModuleRespDtoFunctionType
	ON_PREMISE_INTEGRATION EdgeModuleRespDtoFunctionType
}

func GetEdgeModuleRespDtoFunctionTypeEnum() EdgeModuleRespDtoFunctionTypeEnum {
	return EdgeModuleRespDtoFunctionTypeEnum{
		DATA_PROCESSING: EdgeModuleRespDtoFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: EdgeModuleRespDtoFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: EdgeModuleRespDtoFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
	}
}

func (c EdgeModuleRespDtoFunctionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *EdgeModuleRespDtoFunctionType) UnmarshalJSON(b []byte) error {
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
