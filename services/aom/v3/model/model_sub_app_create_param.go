package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubAppCreateParam struct {

	// 子应用名称：字符集长度2-64，仅支持字符集：英文字母、数字、下划线、中划线、点
	Name string `json:"name"`

	// 子应用节点显示名：字符集长度2-64，仅支持字符集：中文字符、英文字母、数字、下划线、中划线、点
	DisplayName *string `json:"display_name,omitempty"`

	// 应用Id、子应用Id,id长度不能超过36位，由大小写字母、数字组成
	ModelId string `json:"model_id"`

	// 应用、子应用，取值：APPLICATION、SUB_APPLICATION，不区分大小写
	ModelType SubAppCreateParamModelType `json:"model_type"`

	// 描述：最大255字符
	Description *string `json:"description,omitempty"`
}

func (o SubAppCreateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubAppCreateParam struct{}"
	}

	return strings.Join([]string{"SubAppCreateParam", string(data)}, " ")
}

type SubAppCreateParamModelType struct {
	value string
}

type SubAppCreateParamModelTypeEnum struct {
	APPLICATION     SubAppCreateParamModelType
	SUB_APPLICATION SubAppCreateParamModelType
}

func GetSubAppCreateParamModelTypeEnum() SubAppCreateParamModelTypeEnum {
	return SubAppCreateParamModelTypeEnum{
		APPLICATION: SubAppCreateParamModelType{
			value: "APPLICATION",
		},
		SUB_APPLICATION: SubAppCreateParamModelType{
			value: "SUB_APPLICATION",
		},
	}
}

func (c SubAppCreateParamModelType) Value() string {
	return c.value
}

func (c SubAppCreateParamModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubAppCreateParamModelType) UnmarshalJSON(b []byte) error {
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
