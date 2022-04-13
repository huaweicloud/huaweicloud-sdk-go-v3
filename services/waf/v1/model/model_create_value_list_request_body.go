package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建或更新引用表
type CreateValueListRequestBody struct {
	// 引用表名称，只能由数字、字母、中划线、下划线和英文句点组成，长度不能超过64

	Name string `json:"name"`
	// 引用表类型，参见枚举列表

	Type CreateValueListRequestBodyType `json:"type"`
	// 引用表的值

	Values []string `json:"values"`
	// 引用表描述，最长128字符

	Description *string `json:"description,omitempty"`
}

func (o CreateValueListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateValueListRequestBody struct{}"
	}

	return strings.Join([]string{"CreateValueListRequestBody", string(data)}, " ")
}

type CreateValueListRequestBodyType struct {
	value string
}

type CreateValueListRequestBodyTypeEnum struct {
	URL             CreateValueListRequestBodyType
	PARAMS          CreateValueListRequestBodyType
	IP              CreateValueListRequestBodyType
	COOKIE          CreateValueListRequestBodyType
	REFERER         CreateValueListRequestBodyType
	USER_AGENT      CreateValueListRequestBodyType
	HEADER          CreateValueListRequestBodyType
	RESPONSE_CODE   CreateValueListRequestBodyType
	RESPONSE_HEADER CreateValueListRequestBodyType
	RESPONSE_BODY   CreateValueListRequestBodyType
}

func GetCreateValueListRequestBodyTypeEnum() CreateValueListRequestBodyTypeEnum {
	return CreateValueListRequestBodyTypeEnum{
		URL: CreateValueListRequestBodyType{
			value: "url",
		},
		PARAMS: CreateValueListRequestBodyType{
			value: "params",
		},
		IP: CreateValueListRequestBodyType{
			value: "ip",
		},
		COOKIE: CreateValueListRequestBodyType{
			value: "cookie",
		},
		REFERER: CreateValueListRequestBodyType{
			value: "referer",
		},
		USER_AGENT: CreateValueListRequestBodyType{
			value: "user-agent",
		},
		HEADER: CreateValueListRequestBodyType{
			value: "header",
		},
		RESPONSE_CODE: CreateValueListRequestBodyType{
			value: "response_code",
		},
		RESPONSE_HEADER: CreateValueListRequestBodyType{
			value: "response_header",
		},
		RESPONSE_BODY: CreateValueListRequestBodyType{
			value: "response_body",
		},
	}
}

func (c CreateValueListRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateValueListRequestBodyType) UnmarshalJSON(b []byte) error {
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
