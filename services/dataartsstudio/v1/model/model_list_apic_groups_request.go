package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApicGroupsRequest Request Object
type ListApicGroupsRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ListApicGroupsRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 网关实例编号
	ApigInstanceId string `json:"apig_instance_id"`

	// 网关类型
	ApigType ListApicGroupsRequestApigType `json:"apig_type"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
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
	APIGW     ListApicGroupsRequestApigType
	ROMA_APIC ListApicGroupsRequestApigType
}

func GetListApicGroupsRequestApigTypeEnum() ListApicGroupsRequestApigTypeEnum {
	return ListApicGroupsRequestApigTypeEnum{
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
