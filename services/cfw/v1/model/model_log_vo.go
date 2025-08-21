package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogVo struct {

	// **参数解释**： 应用 **取值范围**： 不涉及
	App *string `json:"app,omitempty"`

	// **参数解释**： 流字节数，流量日志字段 **取值范围**： 不涉及
	Bytes *float64 `json:"bytes,omitempty"`

	// **参数解释**： 会话方向 **取值范围**： out2in 外到内访问 in2out 内到外访问
	Direction *string `json:"direction,omitempty"`

	// **参数解释**： 流字节数，访问控制日志和流量日志中存在 **取值范围**： 目的网址
	DstHost *string `json:"dst_host,omitempty"`

	// **参数解释**： 目的IP **取值范围**： 不涉及
	DstIp *string `json:"dst_ip,omitempty"`

	// **参数解释**： 目的端口 **取值范围**： 不涉及
	DstPort *int32 `json:"dst_port,omitempty"`

	// **参数解释**： 会话结束时间，流量日志字段 **取值范围**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 日志ID，用于分页查询 **取值范围**： 不涉及
	LogId *string `json:"log_id,omitempty"`

	// **参数解释**： 流包数，流量日志字段 **取值范围**： 不涉及
	Packets *float64 `json:"packets,omitempty"`

	// **参数解释**： 协议 **取值范围**： 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 源Ip **取值范围**： 不涉及
	SrcIp *string `json:"src_ip,omitempty"`

	// **参数解释**： 源端口 **取值范围**： 不涉及
	SrcPort *int32 `json:"src_port,omitempty"`

	// **参数解释**： 会话开始时间，流量日志字段 **取值范围**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 目的地区ID **取值范围**： 不涉及
	DstRegionId *interface{} `json:"dst_region_id,omitempty"`

	// **参数解释**： 目的地区名称 **取值范围**： 不涉及
	DstRegionName *string `json:"dst_region_name,omitempty"`

	// **参数解释**： 目的省份ID **取值范围**： 不涉及
	DstProvinceId *string `json:"dst_province_id,omitempty"`

	// **参数解释**： 目的省份名称 **取值范围**： 不涉及
	DstProvinceName *string `json:"dst_province_name,omitempty"`

	// **参数解释**： 目的城市ID **取值范围**： 不涉及
	DstCityId *string `json:"dst_city_id,omitempty"`

	// **参数解释**： 目的城市名称 **取值范围**： 不涉及
	DstCityName *string `json:"dst_city_name,omitempty"`

	// **参数解释**： 源地区ID **取值范围**： 不涉及
	SrcRegionId *string `json:"src_region_id,omitempty"`

	// **参数解释**： 源地区名称 **取值范围**： 不涉及
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// **参数解释**： 源省份ID **取值范围**： 不涉及
	SrcProvinceId *string `json:"src_province_id,omitempty"`

	// **参数解释**： 源省份名称 **取值范围**： 不涉及
	SrcProvinceName *string `json:"src_province_name,omitempty"`

	// **参数解释**： 源城市ID **取值范围**： 不涉及
	SrcCityId *string `json:"src_city_id,omitempty"`

	// **参数解释**： 源城市名称 **取值范围**： 不涉及
	SrcCityName *string `json:"src_city_name,omitempty"`

	// **参数解释**： 虚拟网关ID **取值范围**： 不涉及
	VgwId *string `json:"vgw_id,omitempty"`

	// **参数解释**： sctp验证标签，流量日志字段 **取值范围**： 不涉及
	SctpVerificationTag *int64 `json:"sctp_verification_tag,omitempty"`

	// **参数解释**： sctp握手流，流量日志字段 **取值范围**： 不涉及
	SctpIsHandshakeFlow *string `json:"sctp_is_handshake_flow,omitempty"`

	// **参数解释**： qos规则ID，流量日志、访问控制日志存在 **取值范围**： 不涉及
	QosRuleId *string `json:"qos_rule_id,omitempty"`

	// **参数解释**： qos规则名称，流量日志、访问控制日志存在 **取值范围**： 不涉及
	QosRuleName *string `json:"qos_rule_name,omitempty"`

	// **参数解释**： qos通道ID，流量日志字段 **取值范围**： 不涉及
	QosChannelId *string `json:"qos_channel_id,omitempty"`

	// **参数解释**： qos通道名称，流量日志字段 **取值范围**： 不涉及
	QosChannelName *string `json:"qos_channel_name,omitempty"`

	// **参数解释**： qos丢包数，流量日志字段 **取值范围**： 不涉及
	QosDropPackets *float64 `json:"qos_drop_packets,omitempty"`

	// **参数解释**： qos丢弃字节，流量日志字段 **取值范围**： 不涉及
	QosDropBytes *float64 `json:"qos_drop_bytes,omitempty"`

	// **参数解释**： qos规则类型，流量日志、访问控制日志存在 **取值范围**： 不涉及
	QosRuleType *int32 `json:"qos_rule_type,omitempty"`

	// **参数解释**： qos通道类型，流量日志字段 **取值范围**： 不涉及
	QosChannelType *int32 `json:"qos_channel_type,omitempty"`

	// **参数解释**： 动作，攻击日志、访问控制日志、URL日志存在 **取值范围**： 不涉及
	Action *string `json:"action,omitempty"`

	// **参数解释**： url，URL日志字段 **取值范围**： 不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释**： 命中时间，访问控制日志、URL日志存在 **取值范围**： 不涉及
	HitTime *int64 `json:"hit_time,omitempty"`

	// **参数解释**： 规则ID，访问控制日志、URL日志存在 **取值范围**： 不涉及
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： 规则名称，访问控制日志、URL日志存在 **取值范围**： 不涉及
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**： 规则类型，访问控制日志和URL日志字段 **取值范围**： 不涉及
	RuleType *int32 `json:"rule_type,omitempty"`

	// **参数解释**： 规则类型，攻击日志字段 **取值范围**： 不涉及
	AttackRule *string `json:"attack_rule,omitempty"`

	// **参数解释**： 攻击规则ID，攻击日志字段 **取值范围**： 不涉及
	AttackRuleId *string `json:"attack_rule_id,omitempty"`

	// **参数解释**： 攻击类型，攻击日志字段 **取值范围**： 不涉及
	AttackType *string `json:"attack_type,omitempty"`

	// **参数解释**： 事件时间，攻击日志字段 **取值范围**： 不涉及
	EventTime *int64 `json:"event_time,omitempty"`

	// **参数解释**： 攻击等级，攻击日志字段 **取值范围**： 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释**： 规则载荷，攻击日志字段 **取值范围**： 不涉及
	Packet *string `json:"packet,omitempty"`

	// **参数解释**： 攻击来源，攻击日志字段 **取值范围**： 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释**： 真实IP，攻击日志字段 **取值范围**： 不涉及
	RealIp *string `json:"real_ip,omitempty"`

	// **参数解释**： 标签类型，攻击日志字段 **取值范围**： 1 WAF回源IP
	Tag *int32 `json:"tag,omitempty"`
}

func (o LogVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogVo struct{}"
	}

	return strings.Join([]string{"LogVo", string(data)}, " ")
}
