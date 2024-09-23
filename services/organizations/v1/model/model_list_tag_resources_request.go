package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTagResourcesRequest Request Object
type ListTagResourcesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源类型。枚举值：organizations:policies（服务策略）、organizations:ous（组织OU）、organizations:accounts（账号信息） 、organizations:roots：（根）。
	ResourceType ListTagResourcesRequestResourceType `json:"resource_type"`

	// 根、组织单元、账号或策略的唯一标识符（ID）。
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
	ORGANIZATIONSPOLICIES ListTagResourcesRequestResourceType
	ORGANIZATIONSOUS      ListTagResourcesRequestResourceType
	ORGANIZATIONSACCOUNTS ListTagResourcesRequestResourceType
	ORGANIZATIONSROOTS    ListTagResourcesRequestResourceType
}

func GetListTagResourcesRequestResourceTypeEnum() ListTagResourcesRequestResourceTypeEnum {
	return ListTagResourcesRequestResourceTypeEnum{
		ORGANIZATIONSPOLICIES: ListTagResourcesRequestResourceType{
			value: "organizations:policies",
		},
		ORGANIZATIONSOUS: ListTagResourcesRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: ListTagResourcesRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSROOTS: ListTagResourcesRequestResourceType{
			value: "organizations:roots",
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
