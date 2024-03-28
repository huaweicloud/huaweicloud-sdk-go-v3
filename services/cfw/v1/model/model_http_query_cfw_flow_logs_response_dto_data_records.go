package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HttpQueryCfwFlowLogsResponseDtoDataRecords struct {

	// 字节
	Bytes *float64 `json:"bytes,omitempty"`

	// 方向，有内到外和外到内两种
	Direction *HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// 包
	Packets *int32 `json:"packets,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 文档ID
	LogId *string `json:"log_id,omitempty"`

	// 源IP
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口
	SrcPort *int32 `json:"src_port,omitempty"`

	// 目的IP
	DstIp *string `json:"dst_ip,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 目的端口
	DstPort *int32 `json:"dst_port,omitempty"`

	// 协议类型:TCP为6,UDP为17,ICMP为1,ICMPV6为58,ANY为-1,手动类型不为空，自动类型为空
	Protocol *string `json:"protocol,omitempty"`

	// 目标主机
	DstHost *string `json:"dst_host,omitempty"`
}

func (o HttpQueryCfwFlowLogsResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCfwFlowLogsResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"HttpQueryCfwFlowLogsResponseDtoDataRecords", string(data)}, " ")
}

type HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection struct {
	value string
}

type HttpQueryCfwFlowLogsResponseDtoDataRecordsDirectionEnum struct {
	OUT2IN HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection
	IN2OUT HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection
}

func GetHttpQueryCfwFlowLogsResponseDtoDataRecordsDirectionEnum() HttpQueryCfwFlowLogsResponseDtoDataRecordsDirectionEnum {
	return HttpQueryCfwFlowLogsResponseDtoDataRecordsDirectionEnum{
		OUT2IN: HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection{
			value: "out2in",
		},
		IN2OUT: HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection{
			value: "in2out",
		},
	}
}

func (c HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection) Value() string {
	return c.value
}

func (c HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HttpQueryCfwFlowLogsResponseDtoDataRecordsDirection) UnmarshalJSON(b []byte) error {
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
