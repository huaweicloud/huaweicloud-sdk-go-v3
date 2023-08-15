package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateVolumeTagsRequestBody This is a auto create Body Object
type BatchCreateVolumeTagsRequestBody struct {

	// 操作标识，当前支持的取值如下：  添加标签：create
	Action BatchCreateVolumeTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags []Tag `json:"tags"`
}

func (o BatchCreateVolumeTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsRequestBody", string(data)}, " ")
}

type BatchCreateVolumeTagsRequestBodyAction struct {
	value string
}

type BatchCreateVolumeTagsRequestBodyActionEnum struct {
	CREATE BatchCreateVolumeTagsRequestBodyAction
}

func GetBatchCreateVolumeTagsRequestBodyActionEnum() BatchCreateVolumeTagsRequestBodyActionEnum {
	return BatchCreateVolumeTagsRequestBodyActionEnum{
		CREATE: BatchCreateVolumeTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateVolumeTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateVolumeTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateVolumeTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
