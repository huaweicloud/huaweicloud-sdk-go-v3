package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TagResourceV5Request Request Object
type TagResourceV5Request struct {

	// 资源ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ResourceId string `json:"resource_id"`

	// 资源类型，可以为“信任委托”（trust agency）或“IAM用户”（user）。
	ResourceType TagResourceV5RequestResourceType `json:"resource_type"`

	Body *Tags `json:"body,omitempty"`
}

func (o TagResourceV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceV5Request struct{}"
	}

	return strings.Join([]string{"TagResourceV5Request", string(data)}, " ")
}

type TagResourceV5RequestResourceType struct {
	value string
}

type TagResourceV5RequestResourceTypeEnum struct {
	AGENCY TagResourceV5RequestResourceType
	USER   TagResourceV5RequestResourceType
}

func GetTagResourceV5RequestResourceTypeEnum() TagResourceV5RequestResourceTypeEnum {
	return TagResourceV5RequestResourceTypeEnum{
		AGENCY: TagResourceV5RequestResourceType{
			value: "agency",
		},
		USER: TagResourceV5RequestResourceType{
			value: "user",
		},
	}
}

func (c TagResourceV5RequestResourceType) Value() string {
	return c.value
}

func (c TagResourceV5RequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TagResourceV5RequestResourceType) UnmarshalJSON(b []byte) error {
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
