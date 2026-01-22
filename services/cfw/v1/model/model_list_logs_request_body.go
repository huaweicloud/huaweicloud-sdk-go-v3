package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLogsRequestBody **参数解释**： 查询条件 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type ListLogsRequestBody struct {

	// **参数解释**： 过滤条件 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Filters *[]Filter `json:"filters,omitempty"`

	// **参数解释**： 每页显示个数 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`

	// **参数解释**： 偏移量 **约束限制**： 第一页为空，其他页不为空 **取值范围**： 相对于上一页的偏移量 **默认取值**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 文档ID **约束限制**： 第一页为空，其他页不为空 **取值范围**： 上一次查询最后一条数据的log_id **默认取值**： 不涉及
	LogId *string `json:"log_id,omitempty"`

	// **参数解释**： 下个日期 **约束限制**： 第一页为空，其他页不为空 **取值范围**： 查询流量日志时为上一次查询最后一条数据的end_time 查询访问控制日志时为上一次查询最后一条数据的hit_time 查询访问控制日志时为上一次查询最后一条数据的event_time 查询URL日志时为上一次查询最后一条数据的hit_time **默认取值**： 不涉及
	NextDate *int64 `json:"next_date,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime int64 `json:"start_time"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime int64 `json:"end_time"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ListLogsRequestBodyLogType `json:"log_type"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： attack为攻击日志、acl 访问控制日志，flow 流量日志，url url日志 **默认取值**： 不涉及
	Type ListLogsRequestBodyType `json:"type"`
}

func (o ListLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLogsRequestBody", string(data)}, " ")
}

type ListLogsRequestBodyLogType struct {
	value string
}

type ListLogsRequestBodyLogTypeEnum struct {
	INTERNET ListLogsRequestBodyLogType
	NAT      ListLogsRequestBodyLogType
	VPC      ListLogsRequestBodyLogType
	VGW      ListLogsRequestBodyLogType
}

func GetListLogsRequestBodyLogTypeEnum() ListLogsRequestBodyLogTypeEnum {
	return ListLogsRequestBodyLogTypeEnum{
		INTERNET: ListLogsRequestBodyLogType{
			value: "internet",
		},
		NAT: ListLogsRequestBodyLogType{
			value: "nat",
		},
		VPC: ListLogsRequestBodyLogType{
			value: "vpc",
		},
		VGW: ListLogsRequestBodyLogType{
			value: "vgw",
		},
	}
}

func (c ListLogsRequestBodyLogType) Value() string {
	return c.value
}

func (c ListLogsRequestBodyLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLogsRequestBodyLogType) UnmarshalJSON(b []byte) error {
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

type ListLogsRequestBodyType struct {
	value string
}

type ListLogsRequestBodyTypeEnum struct {
	ATTACK ListLogsRequestBodyType
	ACL    ListLogsRequestBodyType
	FLOW   ListLogsRequestBodyType
	URL    ListLogsRequestBodyType
}

func GetListLogsRequestBodyTypeEnum() ListLogsRequestBodyTypeEnum {
	return ListLogsRequestBodyTypeEnum{
		ATTACK: ListLogsRequestBodyType{
			value: "attack",
		},
		ACL: ListLogsRequestBodyType{
			value: "acl",
		},
		FLOW: ListLogsRequestBodyType{
			value: "flow",
		},
		URL: ListLogsRequestBodyType{
			value: "url",
		},
	}
}

func (c ListLogsRequestBodyType) Value() string {
	return c.value
}

func (c ListLogsRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLogsRequestBodyType) UnmarshalJSON(b []byte) error {
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
