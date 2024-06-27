package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HttpQueryCfwAttackLogsResponseDtoDataRecords struct {

	// 方向，有内到外和外到内两种
	Direction *HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// 动作
	Action *string `json:"action,omitempty"`

	// 事件时间，以毫秒为单位的时间戳，如1718936272648
	EventTime *int64 `json:"event_time,omitempty"`

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
	AttackRuleId *string `json:"attack_rule_id,omitempty"`

	// 命中时间，以毫秒为单位的时间戳，如1718936272648
	HitTime *int64 `json:"hit_time,omitempty"`

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

	// 攻击日志报文
	Packet *string `json:"packet,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 攻击报文信息
	PacketMessages *[]PacketMessage `json:"packetMessages,omitempty"`

	// 源区域id
	SrcRegionId *string `json:"src_region_id,omitempty"`

	// 源区域名称
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// 目的区域id
	DstRegionId *string `json:"dst_region_id,omitempty"`

	// 目的区域名称
	DstRegionName *string `json:"dst_region_name,omitempty"`

	// 源省份id
	SrcProvinceId *string `json:"src_province_id,omitempty"`

	// 源省份名称
	SrcProvinceName *string `json:"src_province_name,omitempty"`

	// 源城市id
	SrcCityId *string `json:"src_city_id,omitempty"`

	// 源城市名称
	SrcCityName *string `json:"src_city_name,omitempty"`

	// 目的省份id
	DstProvinceId *string `json:"dst_province_id,omitempty"`

	// 目的省份名称
	DstProvinceName *string `json:"dst_province_name,omitempty"`

	// 目的城市id
	DstCityId *string `json:"dst_city_id,omitempty"`

	// 目的城市名称
	DstCityName *string `json:"dst_city_name,omitempty"`
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
