package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecretRequest Request Object
type CreateSecretRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType CreateSecretRequestContentType `json:"Content-Type"`

	// 项目名称，缺省值默认为区域名称，例如：cn-north-1。
	Projectname *string `json:"projectname,omitempty"`
}

func (o CreateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretRequest", string(data)}, " ")
}

type CreateSecretRequestContentType struct {
	value string
}

type CreateSecretRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateSecretRequestContentType
	APPLICATION_JSON             CreateSecretRequestContentType
}

func GetCreateSecretRequestContentTypeEnum() CreateSecretRequestContentTypeEnum {
	return CreateSecretRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateSecretRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateSecretRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateSecretRequestContentType) Value() string {
	return c.value
}

func (c CreateSecretRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecretRequestContentType) UnmarshalJSON(b []byte) error {
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
