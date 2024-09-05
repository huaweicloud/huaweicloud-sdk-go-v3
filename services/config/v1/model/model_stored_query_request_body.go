package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StoredQueryRequestBody struct {

	// ResourceQL 名字
	Name string `json:"name"`

	// 自定义查询类型，枚举值为“account”和“aggregator”。若取值为“account”，表示单帐号的自定义查询语句；若取值为“aggregator”，表示聚合器的自定义查询语句。默认值为“account”。
	Type *StoredQueryRequestBodyType `json:"type,omitempty"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty"`

	// ResourceQL 表达式
	Expression string `json:"expression"`
}

func (o StoredQueryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoredQueryRequestBody struct{}"
	}

	return strings.Join([]string{"StoredQueryRequestBody", string(data)}, " ")
}

type StoredQueryRequestBodyType struct {
	value string
}

type StoredQueryRequestBodyTypeEnum struct {
	ACCOUNT    StoredQueryRequestBodyType
	AGGREGATOR StoredQueryRequestBodyType
}

func GetStoredQueryRequestBodyTypeEnum() StoredQueryRequestBodyTypeEnum {
	return StoredQueryRequestBodyTypeEnum{
		ACCOUNT: StoredQueryRequestBodyType{
			value: "account",
		},
		AGGREGATOR: StoredQueryRequestBodyType{
			value: "aggregator",
		},
	}
}

func (c StoredQueryRequestBodyType) Value() string {
	return c.value
}

func (c StoredQueryRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StoredQueryRequestBodyType) UnmarshalJSON(b []byte) error {
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
