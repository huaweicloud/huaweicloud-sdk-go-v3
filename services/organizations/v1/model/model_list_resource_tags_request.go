package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceTagsRequest Request Object
type ListResourceTagsRequest struct {

	// 资源类型 organizations:policies服务策略 organizations:ous组织OU organizations:accounts 帐号信息 organizations:roots根
	ResourceType ListResourceTagsRequestResourceType `json:"resource_type"`
}

func (o ListResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceTagsRequest", string(data)}, " ")
}

type ListResourceTagsRequestResourceType struct {
	value string
}

type ListResourceTagsRequestResourceTypeEnum struct {
	ORGANIZATIONSROOTS    ListResourceTagsRequestResourceType
	ORGANIZATIONSOUS      ListResourceTagsRequestResourceType
	ORGANIZATIONSACCOUNTS ListResourceTagsRequestResourceType
	ORGANIZATIONSPOLICIES ListResourceTagsRequestResourceType
}

func GetListResourceTagsRequestResourceTypeEnum() ListResourceTagsRequestResourceTypeEnum {
	return ListResourceTagsRequestResourceTypeEnum{
		ORGANIZATIONSROOTS: ListResourceTagsRequestResourceType{
			value: "organizations:roots",
		},
		ORGANIZATIONSOUS: ListResourceTagsRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: ListResourceTagsRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSPOLICIES: ListResourceTagsRequestResourceType{
			value: "organizations:policies",
		},
	}
}

func (c ListResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
