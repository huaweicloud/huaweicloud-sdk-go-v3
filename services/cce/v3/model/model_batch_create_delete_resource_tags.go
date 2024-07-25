package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateDeleteResourceTags
type BatchCreateDeleteResourceTags struct {

	// action类型，\"create\"或者\"delete\"。
	Action BatchCreateDeleteResourceTagsAction `json:"action"`

	// 资源标签
	Tags *[]ResourceTagBody `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]ResourceTagBody `json:"sys_tags,omitempty"`
}

func (o BatchCreateDeleteResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteResourceTags struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteResourceTags", string(data)}, " ")
}

type BatchCreateDeleteResourceTagsAction struct {
	value string
}

type BatchCreateDeleteResourceTagsActionEnum struct {
	CREATE BatchCreateDeleteResourceTagsAction
	DELETE BatchCreateDeleteResourceTagsAction
}

func GetBatchCreateDeleteResourceTagsActionEnum() BatchCreateDeleteResourceTagsActionEnum {
	return BatchCreateDeleteResourceTagsActionEnum{
		CREATE: BatchCreateDeleteResourceTagsAction{
			value: "create",
		},
		DELETE: BatchCreateDeleteResourceTagsAction{
			value: "delete",
		},
	}
}

func (c BatchCreateDeleteResourceTagsAction) Value() string {
	return c.value
}

func (c BatchCreateDeleteResourceTagsAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateDeleteResourceTagsAction) UnmarshalJSON(b []byte) error {
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
