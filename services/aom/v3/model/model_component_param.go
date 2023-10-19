package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComponentParam struct {

	// 组件描述：最大255字符
	Description *string `json:"description,omitempty"`

	// 应用Id、子应用Id,id长度不能超过36位，由大小写字母、数字组成
	ModelId string `json:"model_id"`

	// 应用、子应用，取值：APPLICATION、SUB_APPLICATION ，不区分大小写
	ModelType ComponentParamModelType `json:"model_type"`

	// 组件名称
	Name string `json:"name"`
}

func (o ComponentParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentParam struct{}"
	}

	return strings.Join([]string{"ComponentParam", string(data)}, " ")
}

type ComponentParamModelType struct {
	value string
}

type ComponentParamModelTypeEnum struct {
	APPLICATION     ComponentParamModelType
	SUB_APPLICATION ComponentParamModelType
}

func GetComponentParamModelTypeEnum() ComponentParamModelTypeEnum {
	return ComponentParamModelTypeEnum{
		APPLICATION: ComponentParamModelType{
			value: "APPLICATION",
		},
		SUB_APPLICATION: ComponentParamModelType{
			value: "SUB_APPLICATION",
		},
	}
}

func (c ComponentParamModelType) Value() string {
	return c.value
}

func (c ComponentParamModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentParamModelType) UnmarshalJSON(b []byte) error {
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
