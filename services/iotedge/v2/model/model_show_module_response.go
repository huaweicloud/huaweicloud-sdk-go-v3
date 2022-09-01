package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowModuleResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty" xml:"edge_app_id"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty" xml:"app_version"`

	// 模块运行状态
	State *ShowModuleResponseState `json:"state,omitempty" xml:"state"`

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
	AppType *ShowModuleResponseAppType `json:"app_type,omitempty" xml:"app_type"`

	// 功能类型
	FunctionType *ShowModuleResponseFunctionType `json:"function_type,omitempty" xml:"function_type"`

	ContainerSettings *ModuleContainerSettingsResDto `json:"container_settings,omitempty" xml:"container_settings"`
	HttpStatusCode    int                            `json:"-"`
}

func (o ShowModuleResponse) String() string {
	data, err := utils.Marshal(o)
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
	DELETE_SUCCESS ShowModuleResponseState
	STOPPED        ShowModuleResponseState
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
		DELETE_SUCCESS: ShowModuleResponseState{
			value: "DELETE_SUCCESS",
		},
		STOPPED: ShowModuleResponseState{
			value: "STOPPED",
		},
	}
}

func (c ShowModuleResponseState) Value() string {
	return c.value
}

func (c ShowModuleResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
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

func (c ShowModuleResponseAppType) Value() string {
	return c.value
}

func (c ShowModuleResponseAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
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
	GATEWAY_MANAGER        ShowModuleResponseFunctionType
	COMPOSITE_APPLICATION  ShowModuleResponseFunctionType
	DATA_COLLECTION        ShowModuleResponseFunctionType
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
		GATEWAY_MANAGER: ShowModuleResponseFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: ShowModuleResponseFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: ShowModuleResponseFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c ShowModuleResponseFunctionType) Value() string {
	return c.value
}

func (c ShowModuleResponseFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
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
