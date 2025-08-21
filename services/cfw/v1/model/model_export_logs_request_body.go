package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportLogsRequestBody **参数解释**： 查询条件 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type ExportLogsRequestBody struct {

	// **参数解释**： 过滤条件 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Filters *[]Filter `json:"filters,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime int64 `json:"start_time"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime int64 `json:"end_time"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ExportLogsRequestBodyLogType `json:"log_type"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： attack为攻击日志、acl 访问控制日志，flow 流量日志，url url日志 **默认取值**： 不涉及
	Type ExportLogsRequestBodyType `json:"type"`

	// **参数解释**： 时区 **约束限制**： 不涉及 **取值范围**： \"GMT+08:00\" **默认取值**： 不涉及
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ExportLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ExportLogsRequestBody", string(data)}, " ")
}

type ExportLogsRequestBodyLogType struct {
	value string
}

type ExportLogsRequestBodyLogTypeEnum struct {
	INTERNET ExportLogsRequestBodyLogType
	NAT      ExportLogsRequestBodyLogType
	VPC      ExportLogsRequestBodyLogType
	VGW      ExportLogsRequestBodyLogType
}

func GetExportLogsRequestBodyLogTypeEnum() ExportLogsRequestBodyLogTypeEnum {
	return ExportLogsRequestBodyLogTypeEnum{
		INTERNET: ExportLogsRequestBodyLogType{
			value: "internet",
		},
		NAT: ExportLogsRequestBodyLogType{
			value: "nat",
		},
		VPC: ExportLogsRequestBodyLogType{
			value: "vpc",
		},
		VGW: ExportLogsRequestBodyLogType{
			value: "vgw",
		},
	}
}

func (c ExportLogsRequestBodyLogType) Value() string {
	return c.value
}

func (c ExportLogsRequestBodyLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportLogsRequestBodyLogType) UnmarshalJSON(b []byte) error {
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

type ExportLogsRequestBodyType struct {
	value string
}

type ExportLogsRequestBodyTypeEnum struct {
	ATTACK ExportLogsRequestBodyType
	ACL    ExportLogsRequestBodyType
	FLOW   ExportLogsRequestBodyType
	URL    ExportLogsRequestBodyType
}

func GetExportLogsRequestBodyTypeEnum() ExportLogsRequestBodyTypeEnum {
	return ExportLogsRequestBodyTypeEnum{
		ATTACK: ExportLogsRequestBodyType{
			value: "attack",
		},
		ACL: ExportLogsRequestBodyType{
			value: "acl",
		},
		FLOW: ExportLogsRequestBodyType{
			value: "flow",
		},
		URL: ExportLogsRequestBodyType{
			value: "url",
		},
	}
}

func (c ExportLogsRequestBodyType) Value() string {
	return c.value
}

func (c ExportLogsRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportLogsRequestBodyType) UnmarshalJSON(b []byte) error {
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
