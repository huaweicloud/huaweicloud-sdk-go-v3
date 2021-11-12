package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建或更新引用表
type CreateValueListRequestBody struct {
	// 引用表名称，2-32位字符串组成

	Name string `json:"name"`
	// 引用表类型，参见枚举列表

	Type CreateValueListRequestBodyType `json:"type"`
	// 引用表的值

	Values *[]string `json:"values,omitempty"`
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
	RESOPNSE_BODY   CreateValueListRequestBodyType
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
		RESOPNSE_BODY: CreateValueListRequestBodyType{
			value: "resopnse_body",
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
