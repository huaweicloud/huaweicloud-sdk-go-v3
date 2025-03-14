package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccessControlLogsRequest Request Object
type ListAccessControlLogsRequest struct {

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// 规则id，可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。
	RuleId *string `json:"rule_id,omitempty"`

	// 开始时间，以毫秒为单位的时间戳，如1718936272648
	StartTime int64 `json:"start_time"`

	// 结束时间，以毫秒为单位的时间戳，如1718936272648
	EndTime int64 `json:"end_time"`

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

	// 规则应用类型包括：“HTTP”，\"HTTPS\"，\"TLS1\"，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”,“BGP”等。
	App *string `json:"app,omitempty"`

	// 文档ID,第一页为空，其他页不为空，其他页可取上一次查询最后一条数据的log_id
	LogId *string `json:"log_id,omitempty"`

	// 下个日期，当是第一页时为空，不是第一页时不为空，其他页可取上一次查询最后一条数据的start_time
	NextDate *int32 `json:"next_date,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于0，首页时为空，非首页时不为空
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 日志类型包括：internet，vpc，nat
	LogType *ListAccessControlLogsRequestLogType `json:"log_type,omitempty"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 目标主机
	DstHost *string `json:"dst_host,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 动作包含permit，deny
	Action *string `json:"action,omitempty"`

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

func (o ListAccessControlLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessControlLogsRequest struct{}"
	}

	return strings.Join([]string{"ListAccessControlLogsRequest", string(data)}, " ")
}

type ListAccessControlLogsRequestLogType struct {
	value string
}

type ListAccessControlLogsRequestLogTypeEnum struct {
	INTERNET ListAccessControlLogsRequestLogType
	NAT      ListAccessControlLogsRequestLogType
	VPC      ListAccessControlLogsRequestLogType
}

func GetListAccessControlLogsRequestLogTypeEnum() ListAccessControlLogsRequestLogTypeEnum {
	return ListAccessControlLogsRequestLogTypeEnum{
		INTERNET: ListAccessControlLogsRequestLogType{
			value: "internet",
		},
		NAT: ListAccessControlLogsRequestLogType{
			value: "nat",
		},
		VPC: ListAccessControlLogsRequestLogType{
			value: "vpc",
		},
	}
}

func (c ListAccessControlLogsRequestLogType) Value() string {
	return c.value
}

func (c ListAccessControlLogsRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccessControlLogsRequestLogType) UnmarshalJSON(b []byte) error {
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
