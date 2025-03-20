package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteResourceTagsV5Request Request Object
type DeleteResourceTagsV5Request struct {

	// 资源ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ResourceId string `json:"resource_id"`

	// 资源类型，可以为“信任委托”（agency）或“IAM用户”（user）。
	ResourceType DeleteResourceTagsV5RequestResourceType `json:"resource_type"`

	// 待删除的标签键列表。
	Body *[]string `json:"body,omitempty"`
}

func (o DeleteResourceTagsV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagsV5Request struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagsV5Request", string(data)}, " ")
}

type DeleteResourceTagsV5RequestResourceType struct {
	value string
}

type DeleteResourceTagsV5RequestResourceTypeEnum struct {
	AGENCY DeleteResourceTagsV5RequestResourceType
	USER   DeleteResourceTagsV5RequestResourceType
}

func GetDeleteResourceTagsV5RequestResourceTypeEnum() DeleteResourceTagsV5RequestResourceTypeEnum {
	return DeleteResourceTagsV5RequestResourceTypeEnum{
		AGENCY: DeleteResourceTagsV5RequestResourceType{
			value: "agency",
		},
		USER: DeleteResourceTagsV5RequestResourceType{
			value: "user",
		},
	}
}

func (c DeleteResourceTagsV5RequestResourceType) Value() string {
	return c.value
}

func (c DeleteResourceTagsV5RequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteResourceTagsV5RequestResourceType) UnmarshalJSON(b []byte) error {
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
