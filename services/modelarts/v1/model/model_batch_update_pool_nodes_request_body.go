package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchUpdatePoolNodesRequestBody struct {

	// **参数解释**：需要更新的节点名称列表。 **约束限制**：不涉及。
	NodeNames []string `json:"nodeNames"`

	// **参数解释**：节点更新的类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - openHaRedundant：开启节点的高可用冗余标签 - closeHaRedundant：关闭节点的高可用冗余标签 - createTags：批量添加节点资源标签 - deleteTags：批量删除节点资源标签 **默认取值**：不涉及。
	Action BatchUpdatePoolNodesRequestBodyAction `json:"action"`

	// **参数解释**：高可用冗余标签效果。 **约束限制**：不涉及。 **取值范围**：可选值如下： - NoSchedule：禁止调度 - NoExecute：禁止执行。 **默认取值**：NoSchedule。
	HaRedundantEffect *string `json:"haRedundantEffect,omitempty"`

	Driver *NodeDriver `json:"driver,omitempty"`

	// **参数解释**：需要批量操作的资源标签列表。 **约束限制**：不涉及。
	Tags *[]NodeTag `json:"tags,omitempty"`
}

func (o BatchUpdatePoolNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoolNodesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoolNodesRequestBody", string(data)}, " ")
}

type BatchUpdatePoolNodesRequestBodyAction struct {
	value string
}

type BatchUpdatePoolNodesRequestBodyActionEnum struct {
	TRUE  BatchUpdatePoolNodesRequestBodyAction
	FALSE BatchUpdatePoolNodesRequestBodyAction
}

func GetBatchUpdatePoolNodesRequestBodyActionEnum() BatchUpdatePoolNodesRequestBodyActionEnum {
	return BatchUpdatePoolNodesRequestBodyActionEnum{
		TRUE: BatchUpdatePoolNodesRequestBodyAction{
			value: "true",
		},
		FALSE: BatchUpdatePoolNodesRequestBodyAction{
			value: "false",
		},
	}
}

func (c BatchUpdatePoolNodesRequestBodyAction) Value() string {
	return c.value
}

func (c BatchUpdatePoolNodesRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdatePoolNodesRequestBodyAction) UnmarshalJSON(b []byte) error {
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
