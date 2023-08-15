package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupsForDomainRequest Request Object
type ListGroupsForDomainRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`

	// 查询的用户组来源
	GroupSource *ListGroupsForDomainRequestGroupSource `json:"group_source,omitempty"`

	// 返回的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListGroupsForDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsForDomainRequest struct{}"
	}

	return strings.Join([]string{"ListGroupsForDomainRequest", string(data)}, " ")
}

type ListGroupsForDomainRequestGroupSource struct {
	value string
}

type ListGroupsForDomainRequestGroupSourceEnum struct {
	IAM   ListGroupsForDomainRequestGroupSource
	SAML  ListGroupsForDomainRequestGroupSource
	LDAP  ListGroupsForDomainRequestGroupSource
	LOCAL ListGroupsForDomainRequestGroupSource
}

func GetListGroupsForDomainRequestGroupSourceEnum() ListGroupsForDomainRequestGroupSourceEnum {
	return ListGroupsForDomainRequestGroupSourceEnum{
		IAM: ListGroupsForDomainRequestGroupSource{
			value: "IAM",
		},
		SAML: ListGroupsForDomainRequestGroupSource{
			value: "SAML",
		},
		LDAP: ListGroupsForDomainRequestGroupSource{
			value: "LDAP",
		},
		LOCAL: ListGroupsForDomainRequestGroupSource{
			value: "LOCAL",
		},
	}
}

func (c ListGroupsForDomainRequestGroupSource) Value() string {
	return c.value
}

func (c ListGroupsForDomainRequestGroupSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupsForDomainRequestGroupSource) UnmarshalJSON(b []byte) error {
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
