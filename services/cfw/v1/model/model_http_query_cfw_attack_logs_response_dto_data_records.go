package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HttpQueryCfwAttackLogsResponseDtoDataRecords struct {

	// 方向，有内到外和外到内两种
	Direction *HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// 动作
	Action *string `json:"action,omitempty"`

	// 事件时间
	EventTime *string `json:"event_time,omitempty"`

	// 攻击类型
	AttackType *string `json:"attack_type,omitempty"`

	// 攻击规则
	AttackRule *string `json:"attack_rule,omitempty"`

	// 威胁等级
	Level *string `json:"level,omitempty"`

	// 来源
	Source *string `json:"source,omitempty"`

	// 报文长度
	PacketLength *int64 `json:"packet_length,omitempty"`

	// 攻击规则id
	AttackRuleId *int32 `json:"attack_rule_id,omitempty"`

	// 命中时间
	HitTime *int32 `json:"hit_time,omitempty"`

	// 日志ID
	LogId *string `json:"log_id,omitempty"`

	// 源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口
	SrcPort *int32 `json:"src_port,omitempty"`

	// 目的IP
	DstIp *string `json:"dst_ip,omitempty"`

	// 目的端口
	DstPort *int32 `json:"dst_port,omitempty"`

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	Packet *Packet `json:"packet,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 攻击报文信息
	PacketMessages *[]PacketMessage `json:"packetMessages,omitempty"`
}

func (o HttpQueryCfwAttackLogsResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwAttackLogsResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwAttackLogsResponseDtoDataRecords", string(data)}, " ")
}

type HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection struct {
	value string
}

type HttpQueryCfwAttackLogsResponseDtoDataRecordsDirectionEnum struct {
	OUT2IN HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection
	IN2OUT HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection
}

func GetHttpQueryCfwAttackLogsResponseDtoDataRecordsDirectionEnum() HttpQueryCfwAttackLogsResponseDtoDataRecordsDirectionEnum {
	return HttpQueryCfwAttackLogsResponseDtoDataRecordsDirectionEnum{
		OUT2IN: HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection{
			value: "out2in",
		},
		IN2OUT: HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection{
			value: "in2out",
		},
	}
}

func (c HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection) Value() string {
	return c.value
}

func (c HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection) UnmarshalJSON(b []byte) error {
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
