package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 函数调用信息
type FunctionRef struct {

	// 函数引用名称，需要和外层functions中的name对应
	RefName *string `json:"ref_name,omitempty"`

	// 函数调用模式，目前只支持同步调用
	InvokeMode *FunctionRefInvokeMode `json:"invoke_mode,omitempty"`

	// 函数执行时的入参，支持引用constants中的常量 定义方式：参数路径 | 常量值/常量路径 参数路径指输入参数的JsonPath路径，如$.a.b[0].c 常量值可以为数字类型，字符串类型(需要用单引号括起来)，布尔类型 常量路径为常量的JsonPath路径，但是根节点需要用$CONST表示，示例：$CONST.a.b
	Arguments *interface{} `json:"arguments,omitempty"`
}

func (o FunctionRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionRef struct{}"
	}

	return strings.Join([]string{"FunctionRef", string(data)}, " ")
}

type FunctionRefInvokeMode struct {
	value string
}

type FunctionRefInvokeModeEnum struct {
	SYNCHRONIZE FunctionRefInvokeMode
}

func GetFunctionRefInvokeModeEnum() FunctionRefInvokeModeEnum {
	return FunctionRefInvokeModeEnum{
		SYNCHRONIZE: FunctionRefInvokeMode{
			value: "synchronize",
		},
	}
}

func (c FunctionRefInvokeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionRefInvokeMode) UnmarshalJSON(b []byte) error {
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
