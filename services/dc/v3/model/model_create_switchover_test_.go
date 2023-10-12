package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSwitchoverTest 创建倒换测试记录对象参数
type CreateSwitchoverTest struct {

	// 倒换测试的资源对象ID
	ResourceId string `json:"resource_id"`

	// 倒换测试的资源对象类型
	ResourceType *CreateSwitchoverTestResourceType `json:"resource_type,omitempty"`

	// shutdown, undo_shutdown表示倒换测试操作类型
	Operation CreateSwitchoverTestOperation `json:"operation"`
}

func (o CreateSwitchoverTest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSwitchoverTest struct{}"
	}

	return strings.Join([]string{"CreateSwitchoverTest", string(data)}, " ")
}

type CreateSwitchoverTestResourceType struct {
	value string
}

type CreateSwitchoverTestResourceTypeEnum struct {
	VIRTUAL_INTERFACE CreateSwitchoverTestResourceType
}

func GetCreateSwitchoverTestResourceTypeEnum() CreateSwitchoverTestResourceTypeEnum {
	return CreateSwitchoverTestResourceTypeEnum{
		VIRTUAL_INTERFACE: CreateSwitchoverTestResourceType{
			value: "virtual_interface",
		},
	}
}

func (c CreateSwitchoverTestResourceType) Value() string {
	return c.value
}

func (c CreateSwitchoverTestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSwitchoverTestResourceType) UnmarshalJSON(b []byte) error {
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

type CreateSwitchoverTestOperation struct {
	value string
}

type CreateSwitchoverTestOperationEnum struct {
	SHUTDOWN      CreateSwitchoverTestOperation
	UNDO_SHUTDOWN CreateSwitchoverTestOperation
}

func GetCreateSwitchoverTestOperationEnum() CreateSwitchoverTestOperationEnum {
	return CreateSwitchoverTestOperationEnum{
		SHUTDOWN: CreateSwitchoverTestOperation{
			value: "shutdown",
		},
		UNDO_SHUTDOWN: CreateSwitchoverTestOperation{
			value: "undo_shutdown",
		},
	}
}

func (c CreateSwitchoverTestOperation) Value() string {
	return c.value
}

func (c CreateSwitchoverTestOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSwitchoverTestOperation) UnmarshalJSON(b []byte) error {
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
