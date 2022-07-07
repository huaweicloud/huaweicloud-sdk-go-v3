package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询模块列表模块相关信息
type EdgeModuleDto struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 模块运行状态
	State *EdgeModuleDtoState `json:"state,omitempty"`

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
	AppType *EdgeModuleDtoAppType `json:"app_type,omitempty"`

	// 功能类型
	FunctionType *EdgeModuleDtoFunctionType `json:"function_type,omitempty"`
}

func (o EdgeModuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeModuleDto struct{}"
	}

	return strings.Join([]string{"EdgeModuleDto", string(data)}, " ")
}

type EdgeModuleDtoState struct {
	value string
}

type EdgeModuleDtoStateEnum struct {
	PENDING        EdgeModuleDtoState
	PENDING_DELETE EdgeModuleDtoState
	DELETE_FAILED  EdgeModuleDtoState
	RUNNING        EdgeModuleDtoState
	FAILED         EdgeModuleDtoState
	SUCCEEDED      EdgeModuleDtoState
	UNKNOWN        EdgeModuleDtoState
	DELETE_SUCCESS EdgeModuleDtoState
	STOPPED        EdgeModuleDtoState
}

func GetEdgeModuleDtoStateEnum() EdgeModuleDtoStateEnum {
	return EdgeModuleDtoStateEnum{
		PENDING: EdgeModuleDtoState{
			value: "PENDING",
		},
		PENDING_DELETE: EdgeModuleDtoState{
			value: "PENDING_DELETE",
		},
		DELETE_FAILED: EdgeModuleDtoState{
			value: "DELETE_FAILED",
		},
		RUNNING: EdgeModuleDtoState{
			value: "RUNNING",
		},
		FAILED: EdgeModuleDtoState{
			value: "FAILED",
		},
		SUCCEEDED: EdgeModuleDtoState{
			value: "SUCCEEDED",
		},
		UNKNOWN: EdgeModuleDtoState{
			value: "UNKNOWN",
		},
		DELETE_SUCCESS: EdgeModuleDtoState{
			value: "DELETE_SUCCESS",
		},
		STOPPED: EdgeModuleDtoState{
			value: "STOPPED",
		},
	}
}

func (c EdgeModuleDtoState) Value() string {
	return c.value
}

func (c EdgeModuleDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EdgeModuleDtoState) UnmarshalJSON(b []byte) error {
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

type EdgeModuleDtoAppType struct {
	value string
}

type EdgeModuleDtoAppTypeEnum struct {
	SYSTEM_REQUIRED EdgeModuleDtoAppType
	SYSTEM_OPTIONAL EdgeModuleDtoAppType
	USER            EdgeModuleDtoAppType
}

func GetEdgeModuleDtoAppTypeEnum() EdgeModuleDtoAppTypeEnum {
	return EdgeModuleDtoAppTypeEnum{
		SYSTEM_REQUIRED: EdgeModuleDtoAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: EdgeModuleDtoAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: EdgeModuleDtoAppType{
			value: "USER",
		},
	}
}

func (c EdgeModuleDtoAppType) Value() string {
	return c.value
}

func (c EdgeModuleDtoAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EdgeModuleDtoAppType) UnmarshalJSON(b []byte) error {
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

type EdgeModuleDtoFunctionType struct {
	value string
}

type EdgeModuleDtoFunctionTypeEnum struct {
	DATA_PROCESSING        EdgeModuleDtoFunctionType
	PROTOCOL_PARSING       EdgeModuleDtoFunctionType
	ON_PREMISE_INTEGRATION EdgeModuleDtoFunctionType
	GATEWAY_MANAGER        EdgeModuleDtoFunctionType
	COMPOSITE_APPLICATION  EdgeModuleDtoFunctionType
	DATA_COLLECTION        EdgeModuleDtoFunctionType
}

func GetEdgeModuleDtoFunctionTypeEnum() EdgeModuleDtoFunctionTypeEnum {
	return EdgeModuleDtoFunctionTypeEnum{
		DATA_PROCESSING: EdgeModuleDtoFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: EdgeModuleDtoFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: EdgeModuleDtoFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
		GATEWAY_MANAGER: EdgeModuleDtoFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: EdgeModuleDtoFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: EdgeModuleDtoFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c EdgeModuleDtoFunctionType) Value() string {
	return c.value
}

func (c EdgeModuleDtoFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EdgeModuleDtoFunctionType) UnmarshalJSON(b []byte) error {
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
