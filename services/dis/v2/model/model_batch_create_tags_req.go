package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateTagsReq struct {

	// 操作标识：仅限于create  - create：批量创建
	Action BatchCreateTagsReqAction `json:"action"`

	// 标签列表。
	Tags []Tag `json:"tags"`
}

func (o BatchCreateTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsReq struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsReq", string(data)}, " ")
}

type BatchCreateTagsReqAction struct {
	value string
}

type BatchCreateTagsReqActionEnum struct {
	CREATE BatchCreateTagsReqAction
}

func GetBatchCreateTagsReqActionEnum() BatchCreateTagsReqActionEnum {
	return BatchCreateTagsReqActionEnum{
		CREATE: BatchCreateTagsReqAction{
			value: "create",
		},
	}
}

func (c BatchCreateTagsReqAction) Value() string {
	return c.value
}

func (c BatchCreateTagsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateTagsReqAction) UnmarshalJSON(b []byte) error {
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
