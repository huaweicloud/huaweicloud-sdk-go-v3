package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KafkaConnectionDetail struct {

	// kafka实例id。
	InstanceId string `json:"instance_id"`

	// kafka连接地址。
	Addr string `json:"addr"`

	// kafka实例是否开启了SASL_SSL。
	SaslSsl bool `json:"sasl_ssl"`

	// kafka实例用户名。实例开启了SASL_SSL时必填
	Username *string `json:"username,omitempty"`

	// kafka实例密码。实例开启了SASL_SSL时必填
	Password *string `json:"password,omitempty"`

	// 收到Server端确认信号个数，表示procuder需要收到多少个这样的确认信号，算消息发送成功。acks参数代表了数据备份的可用性。支持选项： acks=0：表示producer不需要等待任何确认收到的信息，副本将立即加到socket buffer并认为已经发送。没有任何保障可以保证此种情况下server已经成功接收数据，同时重试配置不会发生作用（因为客户端不知道是否失败）回馈的offset会总是设置为-1。 acks=1：这意味着至少要等待leader已经成功将数据写入本地log，但是并没有等待所有follower是否成功写入。如果follower没有成功备份数据，而此时leader又无法提供服务，则消息会丢失。 acks=all：这意味着leader需要等待ISR中所有备份都成功写入日志，只有任何一个备份存活，数据都不会丢失。min.insync.replicas指定必须确认写入才能被认为成功的副本的最小数量。
	Acks *KafkaConnectionDetailAcks `json:"acks,omitempty"`
}

func (o KafkaConnectionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaConnectionDetail struct{}"
	}

	return strings.Join([]string{"KafkaConnectionDetail", string(data)}, " ")
}

type KafkaConnectionDetailAcks struct {
	value string
}

type KafkaConnectionDetailAcksEnum struct {
	E_0 KafkaConnectionDetailAcks
	E_1 KafkaConnectionDetailAcks
	ALL KafkaConnectionDetailAcks
}

func GetKafkaConnectionDetailAcksEnum() KafkaConnectionDetailAcksEnum {
	return KafkaConnectionDetailAcksEnum{
		E_0: KafkaConnectionDetailAcks{
			value: "0",
		},
		E_1: KafkaConnectionDetailAcks{
			value: "1",
		},
		ALL: KafkaConnectionDetailAcks{
			value: "all",
		},
	}
}

func (c KafkaConnectionDetailAcks) Value() string {
	return c.value
}

func (c KafkaConnectionDetailAcks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KafkaConnectionDetailAcks) UnmarshalJSON(b []byte) error {
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
