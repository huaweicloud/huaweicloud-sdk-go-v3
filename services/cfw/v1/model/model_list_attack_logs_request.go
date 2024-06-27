package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAttackLogsRequest Request Object
type ListAttackLogsRequest struct {

	// 开始时间，以毫秒为单位的时间戳，如1718936272648
	StartTime int64 `json:"start_time"`

	// 结束时间，以毫秒为单位的时间戳，如1718936272648
	EndTime int64 `json:"end_time"`

	// 源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口号
	SrcPort *int32 `json:"src_port,omitempty"`

	// 目的IP
	DstIp *string `json:"dst_ip,omitempty"`

	// 目的端口号
	DstPort *int32 `json:"dst_port,omitempty"`

	// 协议类型，包含TCP, UDP,ICMP,ICMPV6等。
	Protocol *string `json:"protocol,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 文档ID,第一页为空，其他页不为空，其他页可取上一次查询最后一条数据的log_id
	LogId *string `json:"log_id,omitempty"`

	// 下个日期，当是第一页时为空，不是第一页时不为空，其他页可取上一次查询最后一条数据的event_time
	NextDate *int64 `json:"next_date,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于0，首页时为空，非首页时不为空
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)。
	FwInstanceId string `json:"fw_instance_id"`

	// 动作包含permit，deny
	Action *string `json:"action,omitempty"`

	// 方向，包含in2out，out2in
	Direction *string `json:"direction,omitempty"`

	// 入侵事件类型
	AttackType *string `json:"attack_type,omitempty"`

	// 入侵事件规则
	AttackRule *string `json:"attack_rule,omitempty"`

	// 威胁等级，包括CRITICAL、HIGH、MEDIUM、LOW
	Level *string `json:"level,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 目标主机
	DstHost *string `json:"dst_host,omitempty"`

	// 日志类型包括：internet，vpc，nat
	LogType *ListAttackLogsRequestLogType `json:"log_type,omitempty"`

	// 入侵事件id
	AttackRuleId *string `json:"attack_rule_id,omitempty"`

	// 源region名称
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// 目的region名称
	DstRegionName *string `json:"dst_region_name,omitempty"`

	// 源省份名称
	SrcProvinceName *string `json:"src_province_name,omitempty"`

	// 目的省份名称
	DstProvinceName *string `json:"dst_province_name,omitempty"`

	// 源城市名称
	SrcCityName *string `json:"src_city_name,omitempty"`

	// 目的城市名称
	DstCityName *string `json:"dst_city_name,omitempty"`
}

func (o ListAttackLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttackLogsRequest struct{}"
	}

	return strings.Join([]string{"ListAttackLogsRequest", string(data)}, " ")
}

type ListAttackLogsRequestLogType struct {
	value string
}

type ListAttackLogsRequestLogTypeEnum struct {
	INTERNET ListAttackLogsRequestLogType
	NAT      ListAttackLogsRequestLogType
	VPC      ListAttackLogsRequestLogType
}

func GetListAttackLogsRequestLogTypeEnum() ListAttackLogsRequestLogTypeEnum {
	return ListAttackLogsRequestLogTypeEnum{
		INTERNET: ListAttackLogsRequestLogType{
			value: "internet",
		},
		NAT: ListAttackLogsRequestLogType{
			value: "nat",
		},
		VPC: ListAttackLogsRequestLogType{
			value: "vpc",
		},
	}
}

func (c ListAttackLogsRequestLogType) Value() string {
	return c.value
}

func (c ListAttackLogsRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackLogsRequestLogType) UnmarshalJSON(b []byte) error {
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
