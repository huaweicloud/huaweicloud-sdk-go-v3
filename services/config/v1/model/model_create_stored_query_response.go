package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateStoredQueryResponse Response Object
type CreateStoredQueryResponse struct {

	// ResourceQL ID
	Id *string `json:"id,omitempty"`

	// ResourceQL 名字
	Name *string `json:"name,omitempty"`

	// 自定义查询类型，枚举值为“account”和“aggregator”。若取值为“account”，表示单帐号的自定义查询语句；若取值为“aggregator”，表示聚合器的自定义查询语句。默认值为“account”。
	Type *CreateStoredQueryResponseType `json:"type,omitempty"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty"`

	// ResourceQL 表达式
	Expression *string `json:"expression,omitempty"`

	// ResourceQL 创建时间
	Created *string `json:"created,omitempty"`

	// ResourceQL 更新时间
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStoredQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStoredQueryResponse struct{}"
	}

	return strings.Join([]string{"CreateStoredQueryResponse", string(data)}, " ")
}

type CreateStoredQueryResponseType struct {
	value string
}

type CreateStoredQueryResponseTypeEnum struct {
	ACCOUNT    CreateStoredQueryResponseType
	AGGREGATOR CreateStoredQueryResponseType
}

func GetCreateStoredQueryResponseTypeEnum() CreateStoredQueryResponseTypeEnum {
	return CreateStoredQueryResponseTypeEnum{
		ACCOUNT: CreateStoredQueryResponseType{
			value: "account",
		},
		AGGREGATOR: CreateStoredQueryResponseType{
			value: "aggregator",
		},
	}
}

func (c CreateStoredQueryResponseType) Value() string {
	return c.value
}

func (c CreateStoredQueryResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateStoredQueryResponseType) UnmarshalJSON(b []byte) error {
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
