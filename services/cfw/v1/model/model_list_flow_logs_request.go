package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlowLogsRequest Request Object
type ListFlowLogsRequest struct {

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。
	FwInstanceId string `json:"fw_instance_id"`

	// 方向
	Direction *string `json:"direction,omitempty"`

	// 日志类型
	LogType *ListFlowLogsRequestLogType `json:"log_type,omitempty"`

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`

	// 源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口
	SrcPort *int32 `json:"src_port,omitempty"`

	// 目的IP
	DstIp *string `json:"dst_ip,omitempty"`

	// 目的端口
	DstPort *int32 `json:"dst_port,omitempty"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *ListFlowLogsRequestProtocol `json:"protocol,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 文档ID，首页时为null，非首页时不为null
	LogId *string `json:"log_id,omitempty"`

	// 日期，首页时为null，非首页时不为null
	NextDate *int64 `json:"next_date,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 目的主机
	DstHost *string `json:"dst_host,omitempty"`
}

func (o ListFlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListFlowLogsRequest", string(data)}, " ")
}

type ListFlowLogsRequestLogType struct {
	value string
}

type ListFlowLogsRequestLogTypeEnum struct {
	INTERNET ListFlowLogsRequestLogType
	VPC      ListFlowLogsRequestLogType
	NAT      ListFlowLogsRequestLogType
}

func GetListFlowLogsRequestLogTypeEnum() ListFlowLogsRequestLogTypeEnum {
	return ListFlowLogsRequestLogTypeEnum{
		INTERNET: ListFlowLogsRequestLogType{
			value: "internet",
		},
		VPC: ListFlowLogsRequestLogType{
			value: "vpc",
		},
		NAT: ListFlowLogsRequestLogType{
			value: "nat",
		},
	}
}

func (c ListFlowLogsRequestLogType) Value() string {
	return c.value
}

func (c ListFlowLogsRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestLogType) UnmarshalJSON(b []byte) error {
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

type ListFlowLogsRequestProtocol struct {
	value string
}

type ListFlowLogsRequestProtocolEnum struct {
	E_6  ListFlowLogsRequestProtocol
	E_17 ListFlowLogsRequestProtocol
	E_1  ListFlowLogsRequestProtocol
	E_58 ListFlowLogsRequestProtocol
}

func GetListFlowLogsRequestProtocolEnum() ListFlowLogsRequestProtocolEnum {
	return ListFlowLogsRequestProtocolEnum{
		E_6: ListFlowLogsRequestProtocol{
			value: "6",
		},
		E_17: ListFlowLogsRequestProtocol{
			value: "17",
		},
		E_1: ListFlowLogsRequestProtocol{
			value: "1",
		},
		E_58: ListFlowLogsRequestProtocol{
			value: "58",
		},
	}
}

func (c ListFlowLogsRequestProtocol) Value() string {
	return c.value
}

func (c ListFlowLogsRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowLogsRequestProtocol) UnmarshalJSON(b []byte) error {
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
