package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceInstancesRequest Request Object
type ListResourceInstancesRequest struct {

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Offset *string `json:"offset,omitempty"`

	// 资源类型 organizations:policies服务策略 organizations:ous组织OU organizations:accounts账号信息 organizations:roots根
	ResourceType ListResourceInstancesRequestResourceType `json:"resource_type"`

	Body *ResourceInstanceReqBody `json:"body,omitempty"`
}

func (o ListResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequest", string(data)}, " ")
}

type ListResourceInstancesRequestResourceType struct {
	value string
}

type ListResourceInstancesRequestResourceTypeEnum struct {
	ORGANIZATIONSPOLICIES ListResourceInstancesRequestResourceType
	ORGANIZATIONSOUS      ListResourceInstancesRequestResourceType
	ORGANIZATIONSACCOUNTS ListResourceInstancesRequestResourceType
	ORGANIZATIONSROOTS    ListResourceInstancesRequestResourceType
}

func GetListResourceInstancesRequestResourceTypeEnum() ListResourceInstancesRequestResourceTypeEnum {
	return ListResourceInstancesRequestResourceTypeEnum{
		ORGANIZATIONSPOLICIES: ListResourceInstancesRequestResourceType{
			value: "organizations:policies",
		},
		ORGANIZATIONSOUS: ListResourceInstancesRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: ListResourceInstancesRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSROOTS: ListResourceInstancesRequestResourceType{
			value: "organizations:roots",
		},
	}
}

func (c ListResourceInstancesRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceInstancesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceInstancesRequestResourceType) UnmarshalJSON(b []byte) error {
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
