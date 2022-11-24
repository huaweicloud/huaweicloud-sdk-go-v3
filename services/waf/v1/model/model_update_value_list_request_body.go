package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建或更新引用表
type UpdateValueListRequestBody struct {

	// 引用表名称，2-32位字符串组成
	Name string `json:"name"`

	// 引用表类型，参见枚举列表
	Type UpdateValueListRequestBodyType `json:"type"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表描述，最长128字符
	Description *string `json:"description,omitempty"`
}

func (o UpdateValueListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateValueListRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateValueListRequestBody", string(data)}, " ")
}

type UpdateValueListRequestBodyType struct {
	value string
}

type UpdateValueListRequestBodyTypeEnum struct {
	URL             UpdateValueListRequestBodyType
	PARAMS          UpdateValueListRequestBodyType
	IP              UpdateValueListRequestBodyType
	COOKIE          UpdateValueListRequestBodyType
	REFERER         UpdateValueListRequestBodyType
	USER_AGENT      UpdateValueListRequestBodyType
	HEADER          UpdateValueListRequestBodyType
	RESPONSE_CODE   UpdateValueListRequestBodyType
	RESPONSE_HEADER UpdateValueListRequestBodyType
	RESOPNSE_BODY   UpdateValueListRequestBodyType
}

func GetUpdateValueListRequestBodyTypeEnum() UpdateValueListRequestBodyTypeEnum {
	return UpdateValueListRequestBodyTypeEnum{
		URL: UpdateValueListRequestBodyType{
			value: "url",
		},
		PARAMS: UpdateValueListRequestBodyType{
			value: "params",
		},
		IP: UpdateValueListRequestBodyType{
			value: "ip",
		},
		COOKIE: UpdateValueListRequestBodyType{
			value: "cookie",
		},
		REFERER: UpdateValueListRequestBodyType{
			value: "referer",
		},
		USER_AGENT: UpdateValueListRequestBodyType{
			value: "user-agent",
		},
		HEADER: UpdateValueListRequestBodyType{
			value: "header",
		},
		RESPONSE_CODE: UpdateValueListRequestBodyType{
			value: "response_code",
		},
		RESPONSE_HEADER: UpdateValueListRequestBodyType{
			value: "response_header",
		},
		RESOPNSE_BODY: UpdateValueListRequestBodyType{
			value: "resopnse_body",
		},
	}
}

func (c UpdateValueListRequestBodyType) Value() string {
	return c.value
}

func (c UpdateValueListRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateValueListRequestBodyType) UnmarshalJSON(b []byte) error {
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
