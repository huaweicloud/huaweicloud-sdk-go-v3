package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeLoadbalancerTagsRequestBody This is a auto create Body Object
type ChangeLoadbalancerTagsRequestBody struct {

	// **参数解释**：操作类型。  **约束限制**：不涉及  **取值范围**： - create：创建标签。 - delete：删除标签。
	Action ChangeLoadbalancerTagsRequestBodyAction `json:"action"`

	// **参数解释**：要操作的标签列表。  **约束限制**：不涉及
	Tags []ChangeResourceTagOption `json:"tags"`
}

func (o ChangeLoadbalancerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerTagsRequestBody", string(data)}, " ")
}

type ChangeLoadbalancerTagsRequestBodyAction struct {
	value string
}

type ChangeLoadbalancerTagsRequestBodyActionEnum struct {
	CREATE ChangeLoadbalancerTagsRequestBodyAction
	DELETE ChangeLoadbalancerTagsRequestBodyAction
}

func GetChangeLoadbalancerTagsRequestBodyActionEnum() ChangeLoadbalancerTagsRequestBodyActionEnum {
	return ChangeLoadbalancerTagsRequestBodyActionEnum{
		CREATE: ChangeLoadbalancerTagsRequestBodyAction{
			value: "create",
		},
		DELETE: ChangeLoadbalancerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c ChangeLoadbalancerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ChangeLoadbalancerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeLoadbalancerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
