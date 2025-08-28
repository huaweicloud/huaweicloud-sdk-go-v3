package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceTagsV5Request Request Object
type ListResourceTagsV5Request struct {

	// 资源ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ResourceId string `json:"resource_id"`

	// 资源类型，可以为“信任委托”（trust agency）或“IAM用户”（user）。
	ResourceType ListResourceTagsV5RequestResourceType `json:"resource_type"`
}

func (o ListResourceTagsV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTagsV5Request struct{}"
	}

	return strings.Join([]string{"ListResourceTagsV5Request", string(data)}, " ")
}

type ListResourceTagsV5RequestResourceType struct {
	value string
}

type ListResourceTagsV5RequestResourceTypeEnum struct {
	AGENCY ListResourceTagsV5RequestResourceType
	USER   ListResourceTagsV5RequestResourceType
}

func GetListResourceTagsV5RequestResourceTypeEnum() ListResourceTagsV5RequestResourceTypeEnum {
	return ListResourceTagsV5RequestResourceTypeEnum{
		AGENCY: ListResourceTagsV5RequestResourceType{
			value: "agency",
		},
		USER: ListResourceTagsV5RequestResourceType{
			value: "user",
		},
	}
}

func (c ListResourceTagsV5RequestResourceType) Value() string {
	return c.value
}

func (c ListResourceTagsV5RequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceTagsV5RequestResourceType) UnmarshalJSON(b []byte) error {
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
