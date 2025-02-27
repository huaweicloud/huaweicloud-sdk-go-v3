package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceConnectionRequest Request Object
type CreateInstanceConnectionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。
	XLanguage *CreateInstanceConnectionRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateInstanceConnectionReq `json:"body,omitempty"`
}

func (o CreateInstanceConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceConnectionRequest", string(data)}, " ")
}

type CreateInstanceConnectionRequestXLanguage struct {
	value string
}

type CreateInstanceConnectionRequestXLanguageEnum struct {
	EN_US CreateInstanceConnectionRequestXLanguage
	ZH_CN CreateInstanceConnectionRequestXLanguage
}

func GetCreateInstanceConnectionRequestXLanguageEnum() CreateInstanceConnectionRequestXLanguageEnum {
	return CreateInstanceConnectionRequestXLanguageEnum{
		EN_US: CreateInstanceConnectionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateInstanceConnectionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateInstanceConnectionRequestXLanguage) Value() string {
	return c.value
}

func (c CreateInstanceConnectionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceConnectionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
