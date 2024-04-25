package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApicGroupsRequest Request Object
type ListApicGroupsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ListApicGroupsRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 网关实例编号，共享版为固定值：APIG。
	ApigInstanceId string `json:"apig_instance_id"`

	// 网关类型。
	ApigType ListApicGroupsRequestApigType `json:"apig_type"`

	// limit。
	Limit *int32 `json:"limit,omitempty"`

	// offset。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListApicGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApicGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListApicGroupsRequest", string(data)}, " ")
}

type ListApicGroupsRequestDlmType struct {
	value string
}

type ListApicGroupsRequestDlmTypeEnum struct {
	SHARED    ListApicGroupsRequestDlmType
	EXCLUSIVE ListApicGroupsRequestDlmType
}

func GetListApicGroupsRequestDlmTypeEnum() ListApicGroupsRequestDlmTypeEnum {
	return ListApicGroupsRequestDlmTypeEnum{
		SHARED: ListApicGroupsRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApicGroupsRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApicGroupsRequestDlmType) Value() string {
	return c.value
}

func (c ListApicGroupsRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApicGroupsRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListApicGroupsRequestApigType struct {
	value string
}

type ListApicGroupsRequestApigTypeEnum struct {
	APIG      ListApicGroupsRequestApigType
	APIGW     ListApicGroupsRequestApigType
	ROMA_APIC ListApicGroupsRequestApigType
}

func GetListApicGroupsRequestApigTypeEnum() ListApicGroupsRequestApigTypeEnum {
	return ListApicGroupsRequestApigTypeEnum{
		APIG: ListApicGroupsRequestApigType{
			value: "APIG",
		},
		APIGW: ListApicGroupsRequestApigType{
			value: "APIGW",
		},
		ROMA_APIC: ListApicGroupsRequestApigType{
			value: "ROMA_APIC",
		},
	}
}

func (c ListApicGroupsRequestApigType) Value() string {
	return c.value
}

func (c ListApicGroupsRequestApigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApicGroupsRequestApigType) UnmarshalJSON(b []byte) error {
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
