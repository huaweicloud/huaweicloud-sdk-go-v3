package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTagResourcesRequest struct {

	// 资源类型 policy服务策略 ou组织OU account帐号信息 root根
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
	POLICY  ListTagResourcesRequestResourceType
	OU      ListTagResourcesRequestResourceType
	ACCOUNT ListTagResourcesRequestResourceType
	ROOT    ListTagResourcesRequestResourceType
}

func GetListTagResourcesRequestResourceTypeEnum() ListTagResourcesRequestResourceTypeEnum {
	return ListTagResourcesRequestResourceTypeEnum{
		POLICY: ListTagResourcesRequestResourceType{
			value: "policy",
		},
		OU: ListTagResourcesRequestResourceType{
			value: "ou",
		},
		ACCOUNT: ListTagResourcesRequestResourceType{
			value: "account",
		},
		ROOT: ListTagResourcesRequestResourceType{
			value: "root",
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
