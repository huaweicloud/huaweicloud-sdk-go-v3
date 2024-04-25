package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAppsRequest Request Object
type ListAppsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListAppsRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// limit。
	Limit *int32 `json:"limit,omitempty"`

	// offset。
	Offset *int32 `json:"offset,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 应用类型。
	AppType *ListAppsRequestAppType `json:"app_type,omitempty"`
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}

type ListAppsRequestDlmType struct {
	value string
}

type ListAppsRequestDlmTypeEnum struct {
	SHARED    ListAppsRequestDlmType
	EXCLUSIVE ListAppsRequestDlmType
}

func GetListAppsRequestDlmTypeEnum() ListAppsRequestDlmTypeEnum {
	return ListAppsRequestDlmTypeEnum{
		SHARED: ListAppsRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListAppsRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListAppsRequestDlmType) Value() string {
	return c.value
}

func (c ListAppsRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppsRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListAppsRequestAppType struct {
	value string
}

type ListAppsRequestAppTypeEnum struct {
	APIG      ListAppsRequestAppType
	IAM       ListAppsRequestAppType
	APIGW     ListAppsRequestAppType
	DLM       ListAppsRequestAppType
	ROMA_APIC ListAppsRequestAppType
}

func GetListAppsRequestAppTypeEnum() ListAppsRequestAppTypeEnum {
	return ListAppsRequestAppTypeEnum{
		APIG: ListAppsRequestAppType{
			value: "APIG",
		},
		IAM: ListAppsRequestAppType{
			value: "IAM",
		},
		APIGW: ListAppsRequestAppType{
			value: "APIGW",
		},
		DLM: ListAppsRequestAppType{
			value: "DLM",
		},
		ROMA_APIC: ListAppsRequestAppType{
			value: "ROMA_APIC",
		},
	}
}

func (c ListAppsRequestAppType) Value() string {
	return c.value
}

func (c ListAppsRequestAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAppsRequestAppType) UnmarshalJSON(b []byte) error {
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
