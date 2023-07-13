package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInterfacesRequest Request Object
type ListInterfacesRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`

	// expression
	Filter *string `json:"filter,omitempty"`

	// 元数据资源全名
	ResourceName *string `json:"resource_name,omitempty"`

	// 元数据资源类型
	ResourceType *ListInterfacesRequestResourceType `json:"resource_type,omitempty"`

	// 授权主体来源
	PrincipalSource *ListInterfacesRequestPrincipalSource `json:"principal_source,omitempty"`

	// 授权主体类型
	PrincipalType *ListInterfacesRequestPrincipalType `json:"principal_type,omitempty"`

	// 授权主体名称
	PrincipalName *string `json:"principal_name,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// page token
	Marker *string `json:"marker,omitempty"`
}

func (o ListInterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterfacesRequest struct{}"
	}

	return strings.Join([]string{"ListInterfacesRequest", string(data)}, " ")
}

type ListInterfacesRequestResourceType struct {
	value string
}

type ListInterfacesRequestResourceTypeEnum struct {
	CATALOG  ListInterfacesRequestResourceType
	DATABASE ListInterfacesRequestResourceType
	TABLE    ListInterfacesRequestResourceType
	FUNC     ListInterfacesRequestResourceType
	MODEL    ListInterfacesRequestResourceType
	COLUMN   ListInterfacesRequestResourceType
	URI      ListInterfacesRequestResourceType
}

func GetListInterfacesRequestResourceTypeEnum() ListInterfacesRequestResourceTypeEnum {
	return ListInterfacesRequestResourceTypeEnum{
		CATALOG: ListInterfacesRequestResourceType{
			value: "CATALOG",
		},
		DATABASE: ListInterfacesRequestResourceType{
			value: "DATABASE",
		},
		TABLE: ListInterfacesRequestResourceType{
			value: "TABLE",
		},
		FUNC: ListInterfacesRequestResourceType{
			value: "FUNC",
		},
		MODEL: ListInterfacesRequestResourceType{
			value: "MODEL",
		},
		COLUMN: ListInterfacesRequestResourceType{
			value: "COLUMN",
		},
		URI: ListInterfacesRequestResourceType{
			value: "URI",
		},
	}
}

func (c ListInterfacesRequestResourceType) Value() string {
	return c.value
}

func (c ListInterfacesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInterfacesRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListInterfacesRequestPrincipalSource struct {
	value string
}

type ListInterfacesRequestPrincipalSourceEnum struct {
	IAM   ListInterfacesRequestPrincipalSource
	SAML  ListInterfacesRequestPrincipalSource
	LDAP  ListInterfacesRequestPrincipalSource
	LOCAL ListInterfacesRequestPrincipalSource
}

func GetListInterfacesRequestPrincipalSourceEnum() ListInterfacesRequestPrincipalSourceEnum {
	return ListInterfacesRequestPrincipalSourceEnum{
		IAM: ListInterfacesRequestPrincipalSource{
			value: "IAM",
		},
		SAML: ListInterfacesRequestPrincipalSource{
			value: "SAML",
		},
		LDAP: ListInterfacesRequestPrincipalSource{
			value: "LDAP",
		},
		LOCAL: ListInterfacesRequestPrincipalSource{
			value: "LOCAL",
		},
	}
}

func (c ListInterfacesRequestPrincipalSource) Value() string {
	return c.value
}

func (c ListInterfacesRequestPrincipalSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInterfacesRequestPrincipalSource) UnmarshalJSON(b []byte) error {
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

type ListInterfacesRequestPrincipalType struct {
	value string
}

type ListInterfacesRequestPrincipalTypeEnum struct {
	USER  ListInterfacesRequestPrincipalType
	GROUP ListInterfacesRequestPrincipalType
	ROLE  ListInterfacesRequestPrincipalType
	SHARE ListInterfacesRequestPrincipalType
	OTHER ListInterfacesRequestPrincipalType
}

func GetListInterfacesRequestPrincipalTypeEnum() ListInterfacesRequestPrincipalTypeEnum {
	return ListInterfacesRequestPrincipalTypeEnum{
		USER: ListInterfacesRequestPrincipalType{
			value: "USER",
		},
		GROUP: ListInterfacesRequestPrincipalType{
			value: "GROUP",
		},
		ROLE: ListInterfacesRequestPrincipalType{
			value: "ROLE",
		},
		SHARE: ListInterfacesRequestPrincipalType{
			value: "SHARE",
		},
		OTHER: ListInterfacesRequestPrincipalType{
			value: "OTHER",
		},
	}
}

func (c ListInterfacesRequestPrincipalType) Value() string {
	return c.value
}

func (c ListInterfacesRequestPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInterfacesRequestPrincipalType) UnmarshalJSON(b []byte) error {
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
