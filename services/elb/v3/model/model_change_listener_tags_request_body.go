package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeListenerTagsRequestBody This is a auto create Body Object
type ChangeListenerTagsRequestBody struct {

	// **参数解释**：操作类型。  **约束限制**：不涉及  **取值范围**： - create：创建标签。 - delete：删除标签。
	Action ChangeListenerTagsRequestBodyAction `json:"action"`

	// **参数解释**：要操作的标签列表。  **约束限制**：不涉及
	Tags []ChangeResourceTagOption `json:"tags"`
}

func (o ChangeListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeListenerTagsRequestBody", string(data)}, " ")
}

type ChangeListenerTagsRequestBodyAction struct {
	value string
}

type ChangeListenerTagsRequestBodyActionEnum struct {
	CREATE ChangeListenerTagsRequestBodyAction
	DELETE ChangeListenerTagsRequestBodyAction
}

func GetChangeListenerTagsRequestBodyActionEnum() ChangeListenerTagsRequestBodyActionEnum {
	return ChangeListenerTagsRequestBodyActionEnum{
		CREATE: ChangeListenerTagsRequestBodyAction{
			value: "create",
		},
		DELETE: ChangeListenerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c ChangeListenerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ChangeListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
