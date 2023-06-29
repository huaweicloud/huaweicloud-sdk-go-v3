package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagResourcesRequest Request Object
type ListTagResourcesRequest struct {

	// 资源类型 organizations:policies服务策略 organizations:ous组织OU organizations:accounts 帐号信息 organizations:roots根
	ResourceType ListTagResourcesRequestResourceType `json:"resource_type"`

	// 根、组织单元、帐号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListTagResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListTagResourcesRequest", string(data)}, " ")
}

type ListTagResourcesRequestResourceType struct {
	value string
}

type ListTagResourcesRequestResourceTypeEnum struct {
	ORGANIZATIONSROOTS    ListTagResourcesRequestResourceType
	ORGANIZATIONSOUS      ListTagResourcesRequestResourceType
	ORGANIZATIONSACCOUNTS ListTagResourcesRequestResourceType
	ORGANIZATIONSPOLICIES ListTagResourcesRequestResourceType
}

func GetListTagResourcesRequestResourceTypeEnum() ListTagResourcesRequestResourceTypeEnum {
	return ListTagResourcesRequestResourceTypeEnum{
		ORGANIZATIONSROOTS: ListTagResourcesRequestResourceType{
			value: "organizations:roots",
		},
		ORGANIZATIONSOUS: ListTagResourcesRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: ListTagResourcesRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSPOLICIES: ListTagResourcesRequestResourceType{
			value: "organizations:policies",
		},
	}
}

func (c ListTagResourcesRequestResourceType) Value() string {
	return c.value
}

func (c ListTagResourcesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagResourcesRequestResourceType) UnmarshalJSON(b []byte) error {
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
