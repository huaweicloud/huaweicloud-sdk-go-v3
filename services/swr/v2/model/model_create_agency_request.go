package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAgencyRequest Request Object
type CreateAgencyRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType CreateAgencyRequestContentType `json:"Content-Type"`
}

func (o CreateAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequest", string(data)}, " ")
}

type CreateAgencyRequestContentType struct {
	value string
}

type CreateAgencyRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateAgencyRequestContentType
	APPLICATION_JSON             CreateAgencyRequestContentType
}

func GetCreateAgencyRequestContentTypeEnum() CreateAgencyRequestContentTypeEnum {
	return CreateAgencyRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateAgencyRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateAgencyRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateAgencyRequestContentType) Value() string {
	return c.value
}

func (c CreateAgencyRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAgencyRequestContentType) UnmarshalJSON(b []byte) error {
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
