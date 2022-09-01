package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListMessageTraceRespTrace struct {

	// 是否成功。
	Success *bool `json:"success,omitempty" xml:"success"`

	// 轨迹类型
	TraceType *ListMessageTraceRespTraceTraceType `json:"trace_type,omitempty" xml:"trace_type"`

	// 时间。
	Timestamp float32 `json:"timestamp,omitempty" xml:"timestamp"`

	// 生产组或消费组。
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 耗时。
	CostTime float32 `json:"cost_time,omitempty" xml:"cost_time"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 消费状态。
	ConsumeStatus float32 `json:"consume_status,omitempty" xml:"consume_status"`

	// 主题名称。
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 消息ID。
	MsgId *string `json:"msg_id,omitempty" xml:"msg_id"`

	// offset消息ID。
	OffsetMsgId *string `json:"offset_msg_id,omitempty" xml:"offset_msg_id"`

	// 消息的标签。
	Tags *string `json:"tags,omitempty" xml:"tags"`

	// 消息的keys。
	Keys *string `json:"keys,omitempty" xml:"keys"`

	// 存储消息的主机IP。
	StoreHost *string `json:"store_host,omitempty" xml:"store_host"`

	// 产生消息的主机IP。
	ClientHost *string `json:"client_host,omitempty" xml:"client_host"`

	// 重试次数。
	RetryTimes *string `json:"retry_times,omitempty" xml:"retry_times"`

	// 消息体长度。
	BodyLength float32 `json:"body_length,omitempty" xml:"body_length"`

	// 消息类型。
	MsgType *ListMessageTraceRespTraceMsgType `json:"msg_type,omitempty" xml:"msg_type"`

	// 事务状态。
	TransactionState *ListMessageTraceRespTraceTransactionState `json:"transaction_state,omitempty" xml:"transaction_state"`

	// 事务ID。
	TransactionId *string `json:"transaction_id,omitempty" xml:"transaction_id"`

	// 是否为事务回查的响应。
	FromTransactionCheck *bool `json:"from_transaction_check,omitempty" xml:"from_transaction_check"`
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
	SUB             ListMessageTraceRespTraceTraceType
	END_TRANSACTION ListMessageTraceRespTraceTraceType
}

func GetListMessageTraceRespTraceTraceTypeEnum() ListMessageTraceRespTraceTraceTypeEnum {
	return ListMessageTraceRespTraceTraceTypeEnum{
		PUB: ListMessageTraceRespTraceTraceType{
			value: "Pub",
		},
		SUB: ListMessageTraceRespTraceTraceType{
			value: "Sub",
		},
		END_TRANSACTION: ListMessageTraceRespTraceTraceType{
			value: "EndTransaction",
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

type ListMessageTraceRespTraceMsgType struct {
	value string
}

type ListMessageTraceRespTraceMsgTypeEnum struct {
	NORMAL_MSG       ListMessageTraceRespTraceMsgType
	TRANS_MSG_HALF   ListMessageTraceRespTraceMsgType
	TRANS_MSG_COMMIT ListMessageTraceRespTraceMsgType
	DELAY_MSG        ListMessageTraceRespTraceMsgType
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
