package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAccessControlLogsRequest struct {

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。
	FwInstanceId string `json:"fw_instance_id"`

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

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

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	// 应用协议
	App *string `json:"app,omitempty"`

	// 文档ID,第一页为空，其他页不为空
	LogId *string `json:"log_id,omitempty"`

	// 日期,第一页为空，其他页不为空
	NextDate *int32 `json:"next_date,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 日志类型
	LogType *ListAccessControlLogsRequestLogType `json:"log_type,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
