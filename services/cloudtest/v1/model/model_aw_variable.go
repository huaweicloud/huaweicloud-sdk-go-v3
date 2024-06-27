package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AwVariable struct {

	// 节点顺序
	ByOrder *int32 `json:"by_order,omitempty"`

	// 当前人员权限
	CurrentPermission *string `json:"currentPermission,omitempty"`

	// 参数描述
	Description *string `json:"description,omitempty"`

	// 动态参数名称
	DynamicParam *string `json:"dynamicParam,omitempty"`

	// 动态参数标志,默认为false:true代表动态参数；false代表非动态参数
	DynamicParamFlag *bool `json:"dynamicParamFlag,omitempty"`

	// 变量参数（[]:空参、[a]:一参，[a,b]:两参）
	FunctionParams *string `json:"functionParams,omitempty"`

	// 响应提取时要使用什么方法处理参数
	FunctionArg *string `json:"function_arg,omitempty"`

	// 响应提取时要使用什么方法处理参数
	FunctionType *AwVariableFunctionType `json:"function_type,omitempty"`

	// 是否是敏感字段
	IsSensitiveInfo *bool `json:"isSensitiveInfo,omitempty"`

	// 敏感字段是否被修改，不敏感字段不关注此值
	IsSensitiveModified *bool `json:"isSensitiveModified,omitempty"`

	// 是否是组合aw的输出参数，只有组合aw下awInstance的aw变量有该字段
	IsOutPut *bool `json:"is_out_put,omitempty"`

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 0/null-变量节点;1-目录节点
	NodeType *int32 `json:"node_type,omitempty"`

	// 旧名称
	OldName *string `json:"oldName,omitempty"`

	// 父节点id
	ParentNodeId *string `json:"parent_node_id,omitempty"`

	// 属性
	Property *string `json:"property,omitempty"`

	// 内置函数枚举
	Regex *string `json:"regex,omitempty"`

	// 敏感参数设置时间
	SensitiveInfoSetterTime *string `json:"sensitiveInfoSetterTime,omitempty"`

	// 敏感参数设置者名称
	SensitiveInfoSetterUser *string `json:"sensitiveInfoSetterUser,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 变量类型（0：text，10-16：7个内置函数）
	VariableType *int32 `json:"variableType,omitempty"`
}

func (o AwVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwVariable struct{}"
	}

	return strings.Join([]string{"AwVariable", string(data)}, " ")
}

type AwVariableFunctionType struct {
	value string
}

type AwVariableFunctionTypeEnum struct {
	REGEX     AwVariableFunctionType
	SUBSTRING AwVariableFunctionType
}

func GetAwVariableFunctionTypeEnum() AwVariableFunctionTypeEnum {
	return AwVariableFunctionTypeEnum{
		REGEX: AwVariableFunctionType{
			value: "REGEX",
		},
		SUBSTRING: AwVariableFunctionType{
			value: "SUBSTRING",
		},
	}
}

func (c AwVariableFunctionType) Value() string {
	return c.value
}

func (c AwVariableFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AwVariableFunctionType) UnmarshalJSON(b []byte) error {
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
