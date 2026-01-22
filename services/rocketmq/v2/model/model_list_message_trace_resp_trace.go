package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListMessageTraceRespTrace struct {

	// **参数解释**： 是否成功。 **约束限制**： 不涉及。 **取值范围**： - true：成功。 - false：失败。 **默认取值**： 不涉及。
	Success *bool `json:"success,omitempty"`

	// **参数解释**： 轨迹类型。 **约束限制**： 不涉及。 **取值范围**： - Pub：生产者成功发送消息。 - SubBefore：消费者准备消费消息。 - SubAfter：消费者完成消息消费。 - EndTransaction：事务消息被提交或回滚。 - Receive：服务侧接收消息。 - Ack：消费者手动确认消费。 **默认取值**： 不涉及。
	TraceType *ListMessageTraceRespTraceTraceType `json:"trace_type,omitempty"`

	// **参数解释**： 时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Timestamp float32 `json:"timestamp,omitempty"`

	// **参数解释**： 生产组或消费组。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 耗时。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CostTime float32 `json:"cost_time,omitempty"`

	// **参数解释**： 请求ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**： 消费状态： **约束限制**： 不涉及。 **取值范围**：  - 0-消费成功  - 1-消费超时  - 2-消费发生异常   - 3-消费返回NULL  - 5-消费失败 **默认取值**： 不涉及。
	ConsumeStatus float32 `json:"consume_status,omitempty"`

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 消息ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	MsgId *string `json:"msg_id,omitempty"`

	// **参数解释**： offset消息ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	OffsetMsgId *string `json:"offset_msg_id,omitempty"`

	// **参数解释**： 消息的标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Tags *string `json:"tags,omitempty"`

	// **参数解释**： 消息的keys。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Keys *string `json:"keys,omitempty"`

	// **参数解释**： 存储消息的主机IP。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	StoreHost *string `json:"store_host,omitempty"`

	// **参数解释**： 产生消息的主机IP。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	ClientHost *string `json:"client_host,omitempty"`

	// **参数解释**： 重试次数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	RetryTimes *int32 `json:"retry_times,omitempty"`

	// **参数解释**： 消息体长度。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	BodyLength float32 `json:"body_length,omitempty"`

	// **参数解释**： 消息类型。 **约束限制**： 不涉及。 **取值范围**： - Normal_Msg：普通消息。 - Trans_Msg_Half：事务半消息。 - Trans_msg_Commit：事务提交消息。 - Delay_Msg：延迟消息。 - Order_Msg：顺序消息。 **默认取值** 不涉及。
	MsgType *ListMessageTraceRespTraceMsgType `json:"msg_type,omitempty"`

	// **参数解释**： 事务状态。 **约束限制**： 不涉及。 **取值范围**： - COMMIT_MESSAGE - ROLLBACK_MESSAGE - UNKNOW **默认取值** 不涉及。
	TransactionState *ListMessageTraceRespTraceTransactionState `json:"transaction_state,omitempty"`

	// **参数解释**： 事务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	TransactionId *string `json:"transaction_id,omitempty"`

	// **参数解释**： 是否为事务回查的响应。 **约束限制**： 不涉及。 **取值范围**： - true：是事务回查的响应。 - false：不是事务回查的响应。 **默认取值** 不涉及。
	FromTransactionCheck *bool `json:"from_transaction_check,omitempty"`
}

func (o ListMessageTraceRespTrace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTraceRespTrace struct{}"
	}

	return strings.Join([]string{"ListMessageTraceRespTrace", string(data)}, " ")
}

type ListMessageTraceRespTraceTraceType struct {
	value string
}

type ListMessageTraceRespTraceTraceTypeEnum struct {
	PUB             ListMessageTraceRespTraceTraceType
	SUB_BEFORE      ListMessageTraceRespTraceTraceType
	SUB_AFTER       ListMessageTraceRespTraceTraceType
	END_TRANSACTION ListMessageTraceRespTraceTraceType
	RECEIVE         ListMessageTraceRespTraceTraceType
	ACK             ListMessageTraceRespTraceTraceType
}

func GetListMessageTraceRespTraceTraceTypeEnum() ListMessageTraceRespTraceTraceTypeEnum {
	return ListMessageTraceRespTraceTraceTypeEnum{
		PUB: ListMessageTraceRespTraceTraceType{
			value: "Pub",
		},
		SUB_BEFORE: ListMessageTraceRespTraceTraceType{
			value: "SubBefore",
		},
		SUB_AFTER: ListMessageTraceRespTraceTraceType{
			value: "SubAfter",
		},
		END_TRANSACTION: ListMessageTraceRespTraceTraceType{
			value: "EndTransaction",
		},
		RECEIVE: ListMessageTraceRespTraceTraceType{
			value: "Receive",
		},
		ACK: ListMessageTraceRespTraceTraceType{
			value: "Ack",
		},
	}
}

func (c ListMessageTraceRespTraceTraceType) Value() string {
	return c.value
}

func (c ListMessageTraceRespTraceTraceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRespTraceTraceType) UnmarshalJSON(b []byte) error {
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

type ListMessageTraceRespTraceMsgType struct {
	value string
}

type ListMessageTraceRespTraceMsgTypeEnum struct {
	NORMAL_MSG       ListMessageTraceRespTraceMsgType
	TRANS_MSG_HALF   ListMessageTraceRespTraceMsgType
	TRANS_MSG_COMMIT ListMessageTraceRespTraceMsgType
	DELAY_MSG        ListMessageTraceRespTraceMsgType
	ORDER_MSG        ListMessageTraceRespTraceMsgType
}

func GetListMessageTraceRespTraceMsgTypeEnum() ListMessageTraceRespTraceMsgTypeEnum {
	return ListMessageTraceRespTraceMsgTypeEnum{
		NORMAL_MSG: ListMessageTraceRespTraceMsgType{
			value: "Normal_Msg",
		},
		TRANS_MSG_HALF: ListMessageTraceRespTraceMsgType{
			value: "Trans_Msg_Half",
		},
		TRANS_MSG_COMMIT: ListMessageTraceRespTraceMsgType{
			value: "Trans_msg_Commit",
		},
		DELAY_MSG: ListMessageTraceRespTraceMsgType{
			value: "Delay_Msg",
		},
		ORDER_MSG: ListMessageTraceRespTraceMsgType{
			value: "Order_Msg",
		},
	}
}

func (c ListMessageTraceRespTraceMsgType) Value() string {
	return c.value
}

func (c ListMessageTraceRespTraceMsgType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRespTraceMsgType) UnmarshalJSON(b []byte) error {
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

type ListMessageTraceRespTraceTransactionState struct {
	value string
}

type ListMessageTraceRespTraceTransactionStateEnum struct {
	COMMIT_MESSAGE   ListMessageTraceRespTraceTransactionState
	ROLLBACK_MESSAGE ListMessageTraceRespTraceTransactionState
	UNKNOW           ListMessageTraceRespTraceTransactionState
}

func GetListMessageTraceRespTraceTransactionStateEnum() ListMessageTraceRespTraceTransactionStateEnum {
	return ListMessageTraceRespTraceTransactionStateEnum{
		COMMIT_MESSAGE: ListMessageTraceRespTraceTransactionState{
			value: "COMMIT_MESSAGE",
		},
		ROLLBACK_MESSAGE: ListMessageTraceRespTraceTransactionState{
			value: "ROLLBACK_MESSAGE",
		},
		UNKNOW: ListMessageTraceRespTraceTransactionState{
			value: "UNKNOW",
		},
	}
}

func (c ListMessageTraceRespTraceTransactionState) Value() string {
	return c.value
}

func (c ListMessageTraceRespTraceTransactionState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRespTraceTransactionState) UnmarshalJSON(b []byte) error {
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
