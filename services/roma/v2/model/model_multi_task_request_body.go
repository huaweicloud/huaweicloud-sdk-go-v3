package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MultiTaskRequestBody struct {

	// 任务名称，只能以字母、数字为开头，包含字母、数字和 . _ -  3~100个字符
	TaskName *string `json:"task_name,omitempty"`

	// 任务ID，可以为空
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型，目前组合任务仅支持实时任务 - REALTIME (实时) - TIMING (定时)
	TaskType *MultiTaskRequestBodyTaskType `json:"task_type,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 任务标签,只能包含字母、数字、中划线、下划线
	TaskTag *string `json:"task_tag,omitempty"`

	// 需要支持的操作类型，支持多选，至少需要选择以下一种： - INSERT - UPDATE - DELETE
	OperationTypes *[]MultiTaskRequestBodyOperationTypes `json:"operation_types,omitempty"`

	// 源端数据源所属集成应用ID
	SourceAppId *string `json:"source_app_id,omitempty"`

	// 目标端数据源所属集成应用ID
	TargetAppId *string `json:"target_app_id,omitempty"`
}

func (o MultiTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskRequestBody struct{}"
	}

	return strings.Join([]string{"MultiTaskRequestBody", string(data)}, " ")
}

type MultiTaskRequestBodyTaskType struct {
	value string
}

type MultiTaskRequestBodyTaskTypeEnum struct {
	REALTIME MultiTaskRequestBodyTaskType
	TIMING   MultiTaskRequestBodyTaskType
}

func GetMultiTaskRequestBodyTaskTypeEnum() MultiTaskRequestBodyTaskTypeEnum {
	return MultiTaskRequestBodyTaskTypeEnum{
		REALTIME: MultiTaskRequestBodyTaskType{
			value: "REALTIME",
		},
		TIMING: MultiTaskRequestBodyTaskType{
			value: "TIMING",
		},
	}
}

func (c MultiTaskRequestBodyTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiTaskRequestBodyTaskType) UnmarshalJSON(b []byte) error {
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

type MultiTaskRequestBodyOperationTypes struct {
	value string
}

type MultiTaskRequestBodyOperationTypesEnum struct {
	INSERT MultiTaskRequestBodyOperationTypes
	UPDATE MultiTaskRequestBodyOperationTypes
	DELETE MultiTaskRequestBodyOperationTypes
}

func GetMultiTaskRequestBodyOperationTypesEnum() MultiTaskRequestBodyOperationTypesEnum {
	return MultiTaskRequestBodyOperationTypesEnum{
		INSERT: MultiTaskRequestBodyOperationTypes{
			value: "INSERT",
		},
		UPDATE: MultiTaskRequestBodyOperationTypes{
			value: "UPDATE",
		},
		DELETE: MultiTaskRequestBodyOperationTypes{
			value: "DELETE",
		},
	}
}

func (c MultiTaskRequestBodyOperationTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiTaskRequestBodyOperationTypes) UnmarshalJSON(b []byte) error {
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
