package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type BatchDeleteVolumeTagsRequestBody struct {
	// 操作标识，当前支持的取值如下：  删除标签：delete

	Action BatchDeleteVolumeTagsRequestBodyAction `json:"action"`
	// 标签列表。

	Tags []DeleteTagsOption `json:"tags"`
}

func (o BatchDeleteVolumeTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsRequestBody", string(data)}, " ")
}

type BatchDeleteVolumeTagsRequestBodyAction struct {
	value string
}

type BatchDeleteVolumeTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteVolumeTagsRequestBodyAction
}

func GetBatchDeleteVolumeTagsRequestBodyActionEnum() BatchDeleteVolumeTagsRequestBodyActionEnum {
	return BatchDeleteVolumeTagsRequestBodyActionEnum{
		DELETE: BatchDeleteVolumeTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteVolumeTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteVolumeTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
