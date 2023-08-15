package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApicInstancesRequest Request Object
type ListApicInstancesRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ListApicInstancesRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 网关类型
	ApigType ListApicInstancesRequestApigType `json:"apig_type"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListApicInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApicInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListApicInstancesRequest", string(data)}, " ")
}

type ListApicInstancesRequestDlmType struct {
	value string
}

type ListApicInstancesRequestDlmTypeEnum struct {
	SHARED    ListApicInstancesRequestDlmType
	EXCLUSIVE ListApicInstancesRequestDlmType
}

func GetListApicInstancesRequestDlmTypeEnum() ListApicInstancesRequestDlmTypeEnum {
	return ListApicInstancesRequestDlmTypeEnum{
		SHARED: ListApicInstancesRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApicInstancesRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApicInstancesRequestDlmType) Value() string {
	return c.value
}

func (c ListApicInstancesRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApicInstancesRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ListApicInstancesRequestApigType struct {
	value string
}

type ListApicInstancesRequestApigTypeEnum struct {
	APIGW     ListApicInstancesRequestApigType
	ROMA_APIC ListApicInstancesRequestApigType
}

func GetListApicInstancesRequestApigTypeEnum() ListApicInstancesRequestApigTypeEnum {
	return ListApicInstancesRequestApigTypeEnum{
		APIGW: ListApicInstancesRequestApigType{
			value: "APIGW",
		},
		ROMA_APIC: ListApicInstancesRequestApigType{
			value: "ROMA_APIC",
		},
	}
}

func (c ListApicInstancesRequestApigType) Value() string {
	return c.value
}

func (c ListApicInstancesRequestApigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApicInstancesRequestApigType) UnmarshalJSON(b []byte) error {
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
