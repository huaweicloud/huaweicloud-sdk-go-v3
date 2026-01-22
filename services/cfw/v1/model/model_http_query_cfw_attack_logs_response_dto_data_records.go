package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HttpQueryCfwAttackLogsResponseDtoDataRecords struct {

	// 方向，包含in2out，out2in
	Direction *HttpQueryCfwAttackLogsResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// 动作包含permit，deny
	Action *string `json:"action,omitempty"`

	// 事件时间，以毫秒为单位的时间戳，如1718936272648
	EventTime *int64 `json:"event_time,omitempty"`

	// 参数解释： 入侵事件类型 约束限制： 不涉及 取值范围： Access Control：访问控制 Vulnerability scanning：漏洞扫描 Email attack：邮件攻击 Vulnerability Attack：漏洞攻击 Web attack：Web攻击 password attack：密码攻击 Hijacking attack：劫持攻击 Protocol exception：协议异常 Trojan horse：特洛伊木马 worms：蠕虫 Buffer Overflow：缓冲区溢出 Hacking tools：黑客工具 Spyware：间谍软件 DDoS flooding：DDoS泛洪 Application-layer DDoS attacks：应用层DDoS攻击 Other suspicious behavior：其他可疑行为 Suspicious DNS activity：可疑DNS活动 Phishing：网络钓鱼 Spam：垃圾邮件 Others：其他攻击 默认取值： 不涉及
	AttackType *string `json:"attack_type,omitempty"`

	// 攻击规则
	AttackRule *string `json:"attack_rule,omitempty"`

	// 威胁等级，包括CRITICAL、HIGH、MEDIUM、LOW
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

	// 协议类型，包含TCP, UDP,ICMP,ICMPV6等。
	Protocol *string `json:"protocol,omitempty"`

	// 攻击日志报文
	Packet *string `json:"packet,omitempty"`

	// 规则应用类型包括：“HTTP”，“HTTPS”，“TLS1”，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”,“BGP”等。
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
