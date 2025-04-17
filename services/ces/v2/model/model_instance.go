package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Instance 资源名称匹配参数
type Instance struct {

	// 云产品名称
	ProductName string `json:"product_name"`

	// 逻辑运算符  ALL 所有条件匹配成功  ANY 任意条件匹配成功
	LogicalOperator InstanceLogicalOperator `json:"logical_operator"`

	// 资源名称匹配参数数组
	InstanceNames []ResourceNameItem `json:"instance_names"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}

type InstanceLogicalOperator struct {
	value string
}

type InstanceLogicalOperatorEnum struct {
	ALL InstanceLogicalOperator
	ANY InstanceLogicalOperator
}

func GetInstanceLogicalOperatorEnum() InstanceLogicalOperatorEnum {
	return InstanceLogicalOperatorEnum{
		ALL: InstanceLogicalOperator{
			value: "ALL",
		},
		ANY: InstanceLogicalOperator{
			value: "ANY",
		},
	}
}

func (c InstanceLogicalOperator) Value() string {
	return c.value
}

func (c InstanceLogicalOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceLogicalOperator) UnmarshalJSON(b []byte) error {
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
