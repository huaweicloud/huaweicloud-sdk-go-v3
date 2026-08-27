package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TxnProgressRequestBody 查询大事务进度的请求体。
type TxnProgressRequestBody struct {

	// **参数解释**： 指定查询的事务动作类型。  **约束限制**：  不涉及。 **取值范围**：  rollback：查询事务的回滚进度。  **默认取值**：   rollback。
	Action TxnProgressRequestBodyAction `json:"action"`

	// **参数解释**：   事务唯一标识列表。   - 列表为空/不传：将执行全量查询，并根据limit和offset分页参数返回当前所有处于执行中的事务信息。   - 列表不为空：将精确匹配并返回transaction_ids中指定的事务信息，此时分页参数（limit/offset）无效。 **约束限制**：   单次查询最多支持100个事务ID。 **取值范围**：   符合事务ID格式的字符串列表。
	TransactionIds *[]string `json:"transaction_ids,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100。  **默认取值**：  100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：    索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。    **约束限制**：    必须为整数，不能为负数。    **取值范围**：    ≥0。  **默认取值**：    0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o TxnProgressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxnProgressRequestBody struct{}"
	}

	return strings.Join([]string{"TxnProgressRequestBody", string(data)}, " ")
}

type TxnProgressRequestBodyAction struct {
	value string
}

type TxnProgressRequestBodyActionEnum struct {
	ROLLBACK TxnProgressRequestBodyAction
}

func GetTxnProgressRequestBodyActionEnum() TxnProgressRequestBodyActionEnum {
	return TxnProgressRequestBodyActionEnum{
		ROLLBACK: TxnProgressRequestBodyAction{
			value: "rollback",
		},
	}
}

func (c TxnProgressRequestBodyAction) Value() string {
	return c.value
}

func (c TxnProgressRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TxnProgressRequestBodyAction) UnmarshalJSON(b []byte) error {
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
