package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FunctionConfig struct {

	// 功能配置id。
	Id *string `json:"id,omitempty"`

	// 功能配置名称。
	Name *string `json:"name,omitempty"`

	// 功能配置开关的状态，表示开启还是关闭 ON/OFF。 - ON： 开启该功能 - OFF： 关闭该功能。
	Status *FunctionConfigStatus `json:"status,omitempty"`

	// 配置项列表，键值对格式。
	Values *[]MapObject `json:"values,omitempty"`
}

func (o FunctionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionConfig struct{}"
	}

	return strings.Join([]string{"FunctionConfig", string(data)}, " ")
}

type FunctionConfigStatus struct {
	value string
}

type FunctionConfigStatusEnum struct {
	ON  FunctionConfigStatus
	OFF FunctionConfigStatus
}

func GetFunctionConfigStatusEnum() FunctionConfigStatusEnum {
	return FunctionConfigStatusEnum{
		ON: FunctionConfigStatus{
			value: "ON",
		},
		OFF: FunctionConfigStatus{
			value: "OFF",
		},
	}
}

func (c FunctionConfigStatus) Value() string {
	return c.value
}

func (c FunctionConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionConfigStatus) UnmarshalJSON(b []byte) error {
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
