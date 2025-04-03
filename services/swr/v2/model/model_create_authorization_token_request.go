package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAuthorizationTokenRequest Request Object
type CreateAuthorizationTokenRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType CreateAuthorizationTokenRequestContentType `json:"Content-Type"`
}

func (o CreateAuthorizationTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizationTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthorizationTokenRequest", string(data)}, " ")
}

type CreateAuthorizationTokenRequestContentType struct {
	value string
}

type CreateAuthorizationTokenRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateAuthorizationTokenRequestContentType
	APPLICATION_JSON             CreateAuthorizationTokenRequestContentType
}

func GetCreateAuthorizationTokenRequestContentTypeEnum() CreateAuthorizationTokenRequestContentTypeEnum {
	return CreateAuthorizationTokenRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateAuthorizationTokenRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateAuthorizationTokenRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateAuthorizationTokenRequestContentType) Value() string {
	return c.value
}

func (c CreateAuthorizationTokenRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAuthorizationTokenRequestContentType) UnmarshalJSON(b []byte) error {
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
