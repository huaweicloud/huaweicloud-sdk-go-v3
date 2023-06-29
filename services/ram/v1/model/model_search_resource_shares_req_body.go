package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchResourceSharesReqBody The request body of the SearchResourceShares operation.
type SearchResourceSharesReqBody struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 资源共享实例名称。
	Name *string `json:"name,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 权限ID。
	PermissionId *string `json:"permission_id,omitempty"`

	// 检索您创建的或共享给您的（self或者other-accounts）资源共享实例。
	ResourceOwner SearchResourceSharesReqBodyResourceOwner `json:"resource_owner"`

	// 资源共享实例的ID列表。
	ResourceShareIds *[]string `json:"resource_share_ids,omitempty"`

	// 资源共享实例的状态。
	ResourceShareStatus *string `json:"resource_share_status,omitempty"`

	// 资源共享实例的标签。
	TagFilters *[]TagFilter `json:"tag_filters,omitempty"`
}

func (o SearchResourceSharesReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceSharesReqBody struct{}"
	}

	return strings.Join([]string{"SearchResourceSharesReqBody", string(data)}, " ")
}

type SearchResourceSharesReqBodyResourceOwner struct {
	value string
}

type SearchResourceSharesReqBodyResourceOwnerEnum struct {
	SELF           SearchResourceSharesReqBodyResourceOwner
	OTHER_ACCOUNTS SearchResourceSharesReqBodyResourceOwner
}

func GetSearchResourceSharesReqBodyResourceOwnerEnum() SearchResourceSharesReqBodyResourceOwnerEnum {
	return SearchResourceSharesReqBodyResourceOwnerEnum{
		SELF: SearchResourceSharesReqBodyResourceOwner{
			value: "self",
		},
		OTHER_ACCOUNTS: SearchResourceSharesReqBodyResourceOwner{
			value: "other-accounts",
		},
	}
}

func (c SearchResourceSharesReqBodyResourceOwner) Value() string {
	return c.value
}

func (c SearchResourceSharesReqBodyResourceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchResourceSharesReqBodyResourceOwner) UnmarshalJSON(b []byte) error {
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
