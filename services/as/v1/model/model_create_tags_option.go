package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTagsOption 标签列表
type CreateTagsOption struct {

	// 标签列表。action为delete时，tags结构体不能缺失，key不能为空，或者空字符串。
	Tags []TagsSingleValue `json:"tags"`

	// 操作标识（区分大小写）：delete：删除。create：创建。若已经存在相同的key值则会覆盖对应的value值。
	Action CreateTagsOptionAction `json:"action"`
}

func (o CreateTagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsOption struct{}"
	}

	return strings.Join([]string{"CreateTagsOption", string(data)}, " ")
}

type CreateTagsOptionAction struct {
	value string
}

type CreateTagsOptionActionEnum struct {
	CREATE CreateTagsOptionAction
}

func GetCreateTagsOptionActionEnum() CreateTagsOptionActionEnum {
	return CreateTagsOptionActionEnum{
		CREATE: CreateTagsOptionAction{
			value: "create",
		},
	}
}

func (c CreateTagsOptionAction) Value() string {
	return c.value
}

func (c CreateTagsOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTagsOptionAction) UnmarshalJSON(b []byte) error {
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
