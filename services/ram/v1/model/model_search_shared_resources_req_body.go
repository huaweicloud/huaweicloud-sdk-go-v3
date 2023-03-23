package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// The request body of the SearchSharedResources operation.
type SearchSharedResourcesReqBody struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 指定资源使用者。
	Principal *string `json:"principal,omitempty"`

	// 指定资源ID列表。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 指定资源URN的列表。
	ResourceUrns *[]string `json:"resource_urns,omitempty"`

	// 指定资源共享实例的所有者（self或者other-accounts）。
	ResourceOwner SearchSharedResourcesReqBodyResourceOwner `json:"resource_owner"`

	// 指定资源共享实例的ID列表。
	ResourceShareIds *[]string `json:"resource_share_ids,omitempty"`

	// 资源所在的区域。
	ResourceRegion *string `json:"resource_region,omitempty"`

	// 指定资源类型。
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o SearchSharedResourcesReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSharedResourcesReqBody struct{}"
	}

	return strings.Join([]string{"SearchSharedResourcesReqBody", string(data)}, " ")
}

type SearchSharedResourcesReqBodyResourceOwner struct {
	value string
}

type SearchSharedResourcesReqBodyResourceOwnerEnum struct {
	SELF           SearchSharedResourcesReqBodyResourceOwner
	OTHER_ACCOUNTS SearchSharedResourcesReqBodyResourceOwner
}

func GetSearchSharedResourcesReqBodyResourceOwnerEnum() SearchSharedResourcesReqBodyResourceOwnerEnum {
	return SearchSharedResourcesReqBodyResourceOwnerEnum{
		SELF: SearchSharedResourcesReqBodyResourceOwner{
			value: "self",
		},
		OTHER_ACCOUNTS: SearchSharedResourcesReqBodyResourceOwner{
			value: "other-accounts",
		},
	}
}

func (c SearchSharedResourcesReqBodyResourceOwner) Value() string {
	return c.value
}

func (c SearchSharedResourcesReqBodyResourceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchSharedResourcesReqBodyResourceOwner) UnmarshalJSON(b []byte) error {
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
