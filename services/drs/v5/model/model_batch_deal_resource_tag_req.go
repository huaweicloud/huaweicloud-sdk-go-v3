package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDealResourceTagReq 批量添加或删除资源标签请求体。
type BatchDealResourceTagReq struct {

	// 操作标识（区分大小写）： - 创建时为“create” - 删除时为“delete”
	Action BatchDealResourceTagReqAction `json:"action"`

	// 标签列表。最多添加10个标签。
	Tags []BatchResourceTag `json:"tags"`
}

func (o BatchDealResourceTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDealResourceTagReq struct{}"
	}

	return strings.Join([]string{"BatchDealResourceTagReq", string(data)}, " ")
}

type BatchDealResourceTagReqAction struct {
	value string
}

type BatchDealResourceTagReqActionEnum struct {
	CREATE BatchDealResourceTagReqAction
	DELETE BatchDealResourceTagReqAction
}

func GetBatchDealResourceTagReqActionEnum() BatchDealResourceTagReqActionEnum {
	return BatchDealResourceTagReqActionEnum{
		CREATE: BatchDealResourceTagReqAction{
			value: "create",
		},
		DELETE: BatchDealResourceTagReqAction{
			value: "delete",
		},
	}
}

func (c BatchDealResourceTagReqAction) Value() string {
	return c.value
}

func (c BatchDealResourceTagReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDealResourceTagReqAction) UnmarshalJSON(b []byte) error {
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
