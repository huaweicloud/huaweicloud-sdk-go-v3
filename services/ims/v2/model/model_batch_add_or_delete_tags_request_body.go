package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// 镜像标签请求体
type BatchAddOrDeleteTagsRequestBody struct {
	// 要进行的标签操作，区分大小写。支持create、delete，分别用于批量地创建/更新、删除标签。
	Action BatchAddOrDeleteTagsRequestBodyAction `json:"action"`
	// 需要增加、修改或者删除的标签键值对集合。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchAddOrDeleteTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsRequestBody", string(data)}, " ")
}

type BatchAddOrDeleteTagsRequestBodyAction struct {
	value string
}

type BatchAddOrDeleteTagsRequestBodyActionEnum struct {
	CREATE BatchAddOrDeleteTagsRequestBodyAction
	DELETE BatchAddOrDeleteTagsRequestBodyAction
}

func GetBatchAddOrDeleteTagsRequestBodyActionEnum() BatchAddOrDeleteTagsRequestBodyActionEnum {
	return BatchAddOrDeleteTagsRequestBodyActionEnum{
		CREATE: BatchAddOrDeleteTagsRequestBodyAction{
			value: "create",
		},
		DELETE: BatchAddOrDeleteTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchAddOrDeleteTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchAddOrDeleteTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
