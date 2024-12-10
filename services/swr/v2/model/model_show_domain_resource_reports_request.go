package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDomainResourceReportsRequest Request Object
type ShowDomainResourceReportsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ShowDomainResourceReportsRequestContentType `json:"Content-Type"`

	// 资源类型
	ResourceType ShowDomainResourceReportsRequestResourceType `json:"resource_type"`

	// 频率类型
	Frequency ShowDomainResourceReportsRequestFrequency `json:"frequency"`
}

func (o ShowDomainResourceReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainResourceReportsRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainResourceReportsRequest", string(data)}, " ")
}

type ShowDomainResourceReportsRequestContentType struct {
	value string
}

type ShowDomainResourceReportsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowDomainResourceReportsRequestContentType
	APPLICATION_JSON             ShowDomainResourceReportsRequestContentType
}

func GetShowDomainResourceReportsRequestContentTypeEnum() ShowDomainResourceReportsRequestContentTypeEnum {
	return ShowDomainResourceReportsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowDomainResourceReportsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowDomainResourceReportsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowDomainResourceReportsRequestContentType) Value() string {
	return c.value
}

func (c ShowDomainResourceReportsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDomainResourceReportsRequestContentType) UnmarshalJSON(b []byte) error {
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

type ShowDomainResourceReportsRequestResourceType struct {
	value string
}

type ShowDomainResourceReportsRequestResourceTypeEnum struct {
	DOWNFLOW ShowDomainResourceReportsRequestResourceType
	STORE    ShowDomainResourceReportsRequestResourceType
}

func GetShowDomainResourceReportsRequestResourceTypeEnum() ShowDomainResourceReportsRequestResourceTypeEnum {
	return ShowDomainResourceReportsRequestResourceTypeEnum{
		DOWNFLOW: ShowDomainResourceReportsRequestResourceType{
			value: "downflow",
		},
		STORE: ShowDomainResourceReportsRequestResourceType{
			value: "store",
		},
	}
}

func (c ShowDomainResourceReportsRequestResourceType) Value() string {
	return c.value
}

func (c ShowDomainResourceReportsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDomainResourceReportsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ShowDomainResourceReportsRequestFrequency struct {
	value string
}

type ShowDomainResourceReportsRequestFrequencyEnum struct {
	DAILY ShowDomainResourceReportsRequestFrequency
}

func GetShowDomainResourceReportsRequestFrequencyEnum() ShowDomainResourceReportsRequestFrequencyEnum {
	return ShowDomainResourceReportsRequestFrequencyEnum{
		DAILY: ShowDomainResourceReportsRequestFrequency{
			value: "daily",
		},
	}
}

func (c ShowDomainResourceReportsRequestFrequency) Value() string {
	return c.value
}

func (c ShowDomainResourceReportsRequestFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDomainResourceReportsRequestFrequency) UnmarshalJSON(b []byte) error {
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
