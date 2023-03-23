package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// The request body of the SearchSharedResources operation.
type SearchDistinctSharedResourcesReqBody struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 指定资源ID。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 指定资源使用者。
	Principal *string `json:"principal,omitempty"`

	// 资源所在的区域。
	ResourceRegion *string `json:"resource_region,omitempty"`

	// 指定资源URN的列表。
	ResourceUrns *[]string `json:"resource_urns,omitempty"`

	// 资源关联的状态。
	Status *string `json:"status,omitempty"`

	// 指定资源共享实例的所有者（self或者other-accounts）。
	ResourceOwner SearchDistinctSharedResourcesReqBodyResourceOwner `json:"resource_owner"`
}

func (o SearchDistinctSharedResourcesReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDistinctSharedResourcesReqBody struct{}"
	}

	return strings.Join([]string{"SearchDistinctSharedResourcesReqBody", string(data)}, " ")
}

type SearchDistinctSharedResourcesReqBodyResourceOwner struct {
	value string
}

type SearchDistinctSharedResourcesReqBodyResourceOwnerEnum struct {
	SELF           SearchDistinctSharedResourcesReqBodyResourceOwner
	OTHER_ACCOUNTS SearchDistinctSharedResourcesReqBodyResourceOwner
}

func GetSearchDistinctSharedResourcesReqBodyResourceOwnerEnum() SearchDistinctSharedResourcesReqBodyResourceOwnerEnum {
	return SearchDistinctSharedResourcesReqBodyResourceOwnerEnum{
		SELF: SearchDistinctSharedResourcesReqBodyResourceOwner{
			value: "self",
		},
		OTHER_ACCOUNTS: SearchDistinctSharedResourcesReqBodyResourceOwner{
			value: "other-accounts",
		},
	}
}

func (c SearchDistinctSharedResourcesReqBodyResourceOwner) Value() string {
	return c.value
}

func (c SearchDistinctSharedResourcesReqBodyResourceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchDistinctSharedResourcesReqBodyResourceOwner) UnmarshalJSON(b []byte) error {
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
