package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// The request body of the SearchSharedPrincipals operation.
type SearchDistinctSharedPrincipalsReqBody struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 指定资源使用者列表。
	Principals *[]string `json:"principals,omitempty"`

	// 指定资源的URN。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 指定资源共享实例的所有者（self或者other-accounts）。
	ResourceOwner SearchDistinctSharedPrincipalsReqBodyResourceOwner `json:"resource_owner"`
}

func (o SearchDistinctSharedPrincipalsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDistinctSharedPrincipalsReqBody struct{}"
	}

	return strings.Join([]string{"SearchDistinctSharedPrincipalsReqBody", string(data)}, " ")
}

type SearchDistinctSharedPrincipalsReqBodyResourceOwner struct {
	value string
}

type SearchDistinctSharedPrincipalsReqBodyResourceOwnerEnum struct {
	SELF           SearchDistinctSharedPrincipalsReqBodyResourceOwner
	OTHER_ACCOUNTS SearchDistinctSharedPrincipalsReqBodyResourceOwner
}

func GetSearchDistinctSharedPrincipalsReqBodyResourceOwnerEnum() SearchDistinctSharedPrincipalsReqBodyResourceOwnerEnum {
	return SearchDistinctSharedPrincipalsReqBodyResourceOwnerEnum{
		SELF: SearchDistinctSharedPrincipalsReqBodyResourceOwner{
			value: "self",
		},
		OTHER_ACCOUNTS: SearchDistinctSharedPrincipalsReqBodyResourceOwner{
			value: "other-accounts",
		},
	}
}

func (c SearchDistinctSharedPrincipalsReqBodyResourceOwner) Value() string {
	return c.value
}

func (c SearchDistinctSharedPrincipalsReqBodyResourceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchDistinctSharedPrincipalsReqBodyResourceOwner) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
