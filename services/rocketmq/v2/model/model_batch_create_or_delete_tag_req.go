package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateOrDeleteTagReq struct {

	// **参数解释**： 操作标识。 **约束限制**： 不涉及。 **取值范围**： - create：创建。 - delete：删除。 **默认取值**： 不涉及。
	Action BatchCreateOrDeleteTagReqAction `json:"action"`

	// **参数解释**： 标签列表。 **约束限制**： 一个RocketMQ实例最多添加20个标签。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags []TagEntity `json:"tags"`
}

func (o BatchCreateOrDeleteTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagReq struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagReq", string(data)}, " ")
}

type BatchCreateOrDeleteTagReqAction struct {
	value string
}

type BatchCreateOrDeleteTagReqActionEnum struct {
	CREATE BatchCreateOrDeleteTagReqAction
	DELETE BatchCreateOrDeleteTagReqAction
}

func GetBatchCreateOrDeleteTagReqActionEnum() BatchCreateOrDeleteTagReqActionEnum {
	return BatchCreateOrDeleteTagReqActionEnum{
		CREATE: BatchCreateOrDeleteTagReqAction{
			value: "create",
		},
		DELETE: BatchCreateOrDeleteTagReqAction{
			value: "delete",
		},
	}
}

func (c BatchCreateOrDeleteTagReqAction) Value() string {
	return c.value
}

func (c BatchCreateOrDeleteTagReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateOrDeleteTagReqAction) UnmarshalJSON(b []byte) error {
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
