package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteTagsReq struct {

	// 操作标识：仅限于delete  - delete：批量删除
	Action BatchDeleteTagsReqAction `json:"action"`

	// 标签列表。
	Tags []Tag `json:"tags"`
}

func (o BatchDeleteTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsReq", string(data)}, " ")
}

type BatchDeleteTagsReqAction struct {
	value string
}

type BatchDeleteTagsReqActionEnum struct {
	DELETE BatchDeleteTagsReqAction
}

func GetBatchDeleteTagsReqActionEnum() BatchDeleteTagsReqActionEnum {
	return BatchDeleteTagsReqActionEnum{
		DELETE: BatchDeleteTagsReqAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteTagsReqAction) Value() string {
	return c.value
}

func (c BatchDeleteTagsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteTagsReqAction) UnmarshalJSON(b []byte) error {
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
