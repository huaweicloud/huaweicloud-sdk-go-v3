package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type MultiTaskUpdateBody struct {
	// 描述信息

	Description *string `json:"description,omitempty"`
	// 任务标签,只能包含字母、数字、中划线、下划线

	TaskTag *string `json:"task_tag,omitempty"`
	// 需要支持的操作类型，支持多选，至少需要选择以下一种： - INSERT - UPDATE - DELETE

	OperationTypes *[]MultiTaskUpdateBodyOperationTypes `json:"operation_types,omitempty"`
	// 是否同步已有数据，仅在编辑任务时生效

	RepullingSnapshot *bool `json:"repulling_snapshot,omitempty"`
}

func (o MultiTaskUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskUpdateBody struct{}"
	}

	return strings.Join([]string{"MultiTaskUpdateBody", string(data)}, " ")
}

type MultiTaskUpdateBodyOperationTypes struct {
	value string
}

type MultiTaskUpdateBodyOperationTypesEnum struct {
	INSERT MultiTaskUpdateBodyOperationTypes
	UPDATE MultiTaskUpdateBodyOperationTypes
	DELETE MultiTaskUpdateBodyOperationTypes
}

func GetMultiTaskUpdateBodyOperationTypesEnum() MultiTaskUpdateBodyOperationTypesEnum {
	return MultiTaskUpdateBodyOperationTypesEnum{
		INSERT: MultiTaskUpdateBodyOperationTypes{
			value: "INSERT",
		},
		UPDATE: MultiTaskUpdateBodyOperationTypes{
			value: "UPDATE",
		},
		DELETE: MultiTaskUpdateBodyOperationTypes{
			value: "DELETE",
		},
	}
}

func (c MultiTaskUpdateBodyOperationTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiTaskUpdateBodyOperationTypes) UnmarshalJSON(b []byte) error {
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
