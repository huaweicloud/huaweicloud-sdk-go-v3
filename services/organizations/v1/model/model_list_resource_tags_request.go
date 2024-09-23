package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceTagsRequest Request Object
type ListResourceTagsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源类型。枚举值：organizations:policies（服务策略）、organizations:ous（组织OU）、organizations:accounts（账号信息） 、organizations:roots：（根）。
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
	ORGANIZATIONSPOLICIES ListResourceTagsRequestResourceType
	ORGANIZATIONSOUS      ListResourceTagsRequestResourceType
	ORGANIZATIONSACCOUNTS ListResourceTagsRequestResourceType
	ORGANIZATIONSROOTS    ListResourceTagsRequestResourceType
}

func GetListResourceTagsRequestResourceTypeEnum() ListResourceTagsRequestResourceTypeEnum {
	return ListResourceTagsRequestResourceTypeEnum{
		ORGANIZATIONSPOLICIES: ListResourceTagsRequestResourceType{
			value: "organizations:policies",
		},
		ORGANIZATIONSOUS: ListResourceTagsRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: ListResourceTagsRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSROOTS: ListResourceTagsRequestResourceType{
			value: "organizations:roots",
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
