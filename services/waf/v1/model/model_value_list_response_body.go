package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValueListResponseBody 单个引用表信息
type ValueListResponseBody struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// 引用表类型
	Type *ValueListResponseBodyType `json:"type,omitempty"`

	// 引用表时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源，1代表用户创建，其它值代表modulleX自动生成
	Producer *int32 `json:"producer,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`
}

func (o ValueListResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueListResponseBody struct{}"
	}

	return strings.Join([]string{"ValueListResponseBody", string(data)}, " ")
}

type ValueListResponseBodyType struct {
	value string
}

type ValueListResponseBodyTypeEnum struct {
	URL             ValueListResponseBodyType
	PARAMS          ValueListResponseBodyType
	IP              ValueListResponseBodyType
	COOKIE          ValueListResponseBodyType
	REFERER         ValueListResponseBodyType
	USER_AGENT      ValueListResponseBodyType
	HEADER          ValueListResponseBodyType
	RESPONSE_CODE   ValueListResponseBodyType
	RESPONSE_HEADER ValueListResponseBodyType
	RESPONSE_BODY   ValueListResponseBodyType
}

func GetValueListResponseBodyTypeEnum() ValueListResponseBodyTypeEnum {
	return ValueListResponseBodyTypeEnum{
		URL: ValueListResponseBodyType{
			value: "url",
		},
		PARAMS: ValueListResponseBodyType{
			value: "params",
		},
		IP: ValueListResponseBodyType{
			value: "ip",
		},
		COOKIE: ValueListResponseBodyType{
			value: "cookie",
		},
		REFERER: ValueListResponseBodyType{
			value: "referer",
		},
		USER_AGENT: ValueListResponseBodyType{
			value: "user-agent",
		},
		HEADER: ValueListResponseBodyType{
			value: "header",
		},
		RESPONSE_CODE: ValueListResponseBodyType{
			value: "response_code",
		},
		RESPONSE_HEADER: ValueListResponseBodyType{
			value: "response_header",
		},
		RESPONSE_BODY: ValueListResponseBodyType{
			value: "response_body",
		},
	}
}

func (c ValueListResponseBodyType) Value() string {
	return c.value
}

func (c ValueListResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValueListResponseBodyType) UnmarshalJSON(b []byte) error {
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
