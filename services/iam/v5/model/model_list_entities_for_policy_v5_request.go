package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEntitiesForPolicyV5Request Request Object
type ListEntitiesForPolicyV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	// 实体类型。
	EntityType *ListEntitiesForPolicyV5RequestEntityType `json:"entity_type,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`
}

func (o ListEntitiesForPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesForPolicyV5Request struct{}"
	}

	return strings.Join([]string{"ListEntitiesForPolicyV5Request", string(data)}, " ")
}

type ListEntitiesForPolicyV5RequestEntityType struct {
	value string
}

type ListEntitiesForPolicyV5RequestEntityTypeEnum struct {
	USER   ListEntitiesForPolicyV5RequestEntityType
	GROUP  ListEntitiesForPolicyV5RequestEntityType
	AGENCY ListEntitiesForPolicyV5RequestEntityType
}

func GetListEntitiesForPolicyV5RequestEntityTypeEnum() ListEntitiesForPolicyV5RequestEntityTypeEnum {
	return ListEntitiesForPolicyV5RequestEntityTypeEnum{
		USER: ListEntitiesForPolicyV5RequestEntityType{
			value: "user",
		},
		GROUP: ListEntitiesForPolicyV5RequestEntityType{
			value: "group",
		},
		AGENCY: ListEntitiesForPolicyV5RequestEntityType{
			value: "agency",
		},
	}
}

func (c ListEntitiesForPolicyV5RequestEntityType) Value() string {
	return c.value
}

func (c ListEntitiesForPolicyV5RequestEntityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEntitiesForPolicyV5RequestEntityType) UnmarshalJSON(b []byte) error {
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
