package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Rule struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty"`

	// 规则ID
	RuleId *int32 `json:"rule_id,omitempty"`

	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64
	Name *string `json:"name,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 规则状态 0-启用 1-停用
	Status *RuleStatus `json:"status,omitempty"`

	// 数据解析状态，ENABLE时data_parsing必填 0-启用 1-停用
	DataParsingStatus *RuleDataParsingStatus `json:"data_parsing_status,omitempty"`

	// SQL查询字段
	SqlField *string `json:"sql_field,omitempty"`

	// SQL查询条件
	SqlWhere *string `json:"sql_where,omitempty"`

	// 完整的规则表达式
	RuleExpress *string `json:"rule_express,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty"`

	// 最后修改时间，timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}

type RuleStatus struct {
	value int32
}

type RuleStatusEnum struct {
	E_0 RuleStatus
	E_1 RuleStatus
}

func GetRuleStatusEnum() RuleStatusEnum {
	return RuleStatusEnum{
		E_0: RuleStatus{
			value: 0,
		}, E_1: RuleStatus{
			value: 1,
		},
	}
}

func (c RuleStatus) Value() int32 {
	return c.value
}

func (c RuleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type RuleDataParsingStatus struct {
	value int32
}

type RuleDataParsingStatusEnum struct {
	E_0 RuleDataParsingStatus
	E_1 RuleDataParsingStatus
}

func GetRuleDataParsingStatusEnum() RuleDataParsingStatusEnum {
	return RuleDataParsingStatusEnum{
		E_0: RuleDataParsingStatus{
			value: 0,
		}, E_1: RuleDataParsingStatus{
			value: 1,
		},
	}
}

func (c RuleDataParsingStatus) Value() int32 {
	return c.value
}

func (c RuleDataParsingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleDataParsingStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
