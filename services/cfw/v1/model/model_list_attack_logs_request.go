package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAttackLogsRequest Request Object
type ListAttackLogsRequest struct {

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`

	// 源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口号
	SrcPort *int32 `json:"src_port,omitempty"`

	// 目的IP
	DstIp *string `json:"dst_ip,omitempty"`

	// 目的端口号
	DstPort *int32 `json:"dst_port,omitempty"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *ListAttackLogsRequestProtocol `json:"protocol,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 日志ID，当是第一页时为空，不是第一页时不为空
	LogId *string `json:"log_id,omitempty"`

	// 下个日期，当是第一页时为空，不是第一页时不为空
	NextDate *int64 `json:"next_date,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。
	FwInstanceId string `json:"fw_instance_id"`

	// 动作0：permit,1：deny
	Action *ListAttackLogsRequestAction `json:"action,omitempty"`

	// 方向0：外到内1：内到外
	Direction *ListAttackLogsRequestDirection `json:"direction,omitempty"`

	// 入侵事件类型
	AttackType *string `json:"attack_type,omitempty"`

	// 入侵事件规则
	AttackRule *string `json:"attack_rule,omitempty"`

	// 威胁等级
	Level *string `json:"level,omitempty"`

	// 判断来源
	Source *string `json:"source,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 目标主机
	DstHost *string `json:"dst_host,omitempty"`

	// 日志类型
	LogType *ListAttackLogsRequestLogType `json:"log_type,omitempty"`

	// 入侵事件id
	AttackRuleId *string `json:"attack_rule_id,omitempty"`

	// 源region名称
	SrcRegionName *string `json:"src_region_name,omitempty"`

	// 目的region名称
	DstRegionName *string `json:"dst_region_name,omitempty"`
}

func (o ListAttackLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttackLogsRequest struct{}"
	}

	return strings.Join([]string{"ListAttackLogsRequest", string(data)}, " ")
}

type ListAttackLogsRequestProtocol struct {
	value string
}

type ListAttackLogsRequestProtocolEnum struct {
	E_6  ListAttackLogsRequestProtocol
	E_17 ListAttackLogsRequestProtocol
	E_1  ListAttackLogsRequestProtocol
	E_58 ListAttackLogsRequestProtocol
}

func GetListAttackLogsRequestProtocolEnum() ListAttackLogsRequestProtocolEnum {
	return ListAttackLogsRequestProtocolEnum{
		E_6: ListAttackLogsRequestProtocol{
			value: "6",
		},
		E_17: ListAttackLogsRequestProtocol{
			value: "17",
		},
		E_1: ListAttackLogsRequestProtocol{
			value: "1",
		},
		E_58: ListAttackLogsRequestProtocol{
			value: "58",
		},
	}
}

func (c ListAttackLogsRequestProtocol) Value() string {
	return c.value
}

func (c ListAttackLogsRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackLogsRequestProtocol) UnmarshalJSON(b []byte) error {
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

type ListAttackLogsRequestAction struct {
	value string
}

type ListAttackLogsRequestActionEnum struct {
	E_0 ListAttackLogsRequestAction
	E_1 ListAttackLogsRequestAction
}

func GetListAttackLogsRequestActionEnum() ListAttackLogsRequestActionEnum {
	return ListAttackLogsRequestActionEnum{
		E_0: ListAttackLogsRequestAction{
			value: "0",
		},
		E_1: ListAttackLogsRequestAction{
			value: "1",
		},
	}
}

func (c ListAttackLogsRequestAction) Value() string {
	return c.value
}

func (c ListAttackLogsRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackLogsRequestAction) UnmarshalJSON(b []byte) error {
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

type ListAttackLogsRequestDirection struct {
	value string
}

type ListAttackLogsRequestDirectionEnum struct {
	E_0 ListAttackLogsRequestDirection
	E_1 ListAttackLogsRequestDirection
}

func GetListAttackLogsRequestDirectionEnum() ListAttackLogsRequestDirectionEnum {
	return ListAttackLogsRequestDirectionEnum{
		E_0: ListAttackLogsRequestDirection{
			value: "0",
		},
		E_1: ListAttackLogsRequestDirection{
			value: "1",
		},
	}
}

func (c ListAttackLogsRequestDirection) Value() string {
	return c.value
}

func (c ListAttackLogsRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackLogsRequestDirection) UnmarshalJSON(b []byte) error {
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
