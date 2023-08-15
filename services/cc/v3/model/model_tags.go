package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Tags 资源标签Tag汇总复数
type Tags struct {

	// 动作。|- create：创建。 delete：删除。
	Action *TagsAction `json:"action,omitempty"`

	// 批量添加/删除资源标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}

type TagsAction struct {
	value string
}

type TagsActionEnum struct {
	CREATE TagsAction
	DELETE TagsAction
}

func GetTagsActionEnum() TagsActionEnum {
	return TagsActionEnum{
		CREATE: TagsAction{
			value: "create",
		},
		DELETE: TagsAction{
			value: "delete",
		},
	}
}

func (c TagsAction) Value() string {
	return c.value
}

func (c TagsAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TagsAction) UnmarshalJSON(b []byte) error {
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
