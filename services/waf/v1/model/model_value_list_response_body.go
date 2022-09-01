package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 单个引用表信息
type ValueListResponseBody struct {

	// 引用表id
	Id *string `json:"id,omitempty" xml:"id"`

	// 引用表名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 引用表类型
	Type *ValueListResponseBodyType `json:"type,omitempty" xml:"type"`

	// 引用表时间戳
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 引用表的值
	Values *[]string `json:"values,omitempty" xml:"values"`

	// 引用表来源，1代表用户创建，其它值代表modulleX自动生成
	Producer *int32 `json:"producer,omitempty" xml:"producer"`

	// 引用表描述
	Description *string `json:"description,omitempty" xml:"description"`
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
