package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListResourceTagsRequest struct {

	// 资源类型 policy服务策略 ou组织OU account帐号信息 root根
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
	POLICY  ListResourceTagsRequestResourceType
	OU      ListResourceTagsRequestResourceType
	ACCOUNT ListResourceTagsRequestResourceType
	ROOT    ListResourceTagsRequestResourceType
}

func GetListResourceTagsRequestResourceTypeEnum() ListResourceTagsRequestResourceTypeEnum {
	return ListResourceTagsRequestResourceTypeEnum{
		POLICY: ListResourceTagsRequestResourceType{
			value: "policy",
		},
		OU: ListResourceTagsRequestResourceType{
			value: "ou",
		},
		ACCOUNT: ListResourceTagsRequestResourceType{
			value: "account",
		},
		ROOT: ListResourceTagsRequestResourceType{
			value: "root",
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
