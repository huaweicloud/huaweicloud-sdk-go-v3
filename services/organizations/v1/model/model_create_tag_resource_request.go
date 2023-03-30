package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateTagResourceRequest struct {

	// 资源类型 organizations:policies服务策略 organizations:ous组织OU organizations:accounts 帐号信息 organizations:roots根
	ResourceType CreateTagResourceRequestResourceType `json:"resource_type"`

	// 根、组织单元、帐号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	Body *TagResourceReqBody `json:"body,omitempty"`
}

func (o CreateTagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateTagResourceRequest", string(data)}, " ")
}

type CreateTagResourceRequestResourceType struct {
	value string
}

type CreateTagResourceRequestResourceTypeEnum struct {
	ORGANIZATIONSROOTS    CreateTagResourceRequestResourceType
	ORGANIZATIONSOUS      CreateTagResourceRequestResourceType
	ORGANIZATIONSACCOUNTS CreateTagResourceRequestResourceType
	ORGANIZATIONSPOLICIES CreateTagResourceRequestResourceType
}

func GetCreateTagResourceRequestResourceTypeEnum() CreateTagResourceRequestResourceTypeEnum {
	return CreateTagResourceRequestResourceTypeEnum{
		ORGANIZATIONSROOTS: CreateTagResourceRequestResourceType{
			value: "organizations:roots",
		},
		ORGANIZATIONSOUS: CreateTagResourceRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: CreateTagResourceRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSPOLICIES: CreateTagResourceRequestResourceType{
			value: "organizations:policies",
		},
	}
}

func (c CreateTagResourceRequestResourceType) Value() string {
	return c.value
}

func (c CreateTagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
