package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchSharedPrincipalsReqBody The request body of the SearchSharedPrincipals operation.
type SearchSharedPrincipalsReqBody struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 指定资源使用者。
	Principals *[]string `json:"principals,omitempty"`

	// 指定资源的URN。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 指定资源共享实例的所有者（self或者other-accounts）。
	ResourceOwner SearchSharedPrincipalsReqBodyResourceOwner `json:"resource_owner"`

	// 指定资源共享实例的ID列表。
	ResourceShareIds *[]string `json:"resource_share_ids,omitempty"`
}

func (o SearchSharedPrincipalsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSharedPrincipalsReqBody struct{}"
	}

	return strings.Join([]string{"SearchSharedPrincipalsReqBody", string(data)}, " ")
}

type SearchSharedPrincipalsReqBodyResourceOwner struct {
	value string
}

type SearchSharedPrincipalsReqBodyResourceOwnerEnum struct {
	SELF           SearchSharedPrincipalsReqBodyResourceOwner
	OTHER_ACCOUNTS SearchSharedPrincipalsReqBodyResourceOwner
}

func GetSearchSharedPrincipalsReqBodyResourceOwnerEnum() SearchSharedPrincipalsReqBodyResourceOwnerEnum {
	return SearchSharedPrincipalsReqBodyResourceOwnerEnum{
		SELF: SearchSharedPrincipalsReqBodyResourceOwner{
			value: "self",
		},
		OTHER_ACCOUNTS: SearchSharedPrincipalsReqBodyResourceOwner{
			value: "other-accounts",
		},
	}
}

func (c SearchSharedPrincipalsReqBodyResourceOwner) Value() string {
	return c.value
}

func (c SearchSharedPrincipalsReqBodyResourceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchSharedPrincipalsReqBodyResourceOwner) UnmarshalJSON(b []byte) error {
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
