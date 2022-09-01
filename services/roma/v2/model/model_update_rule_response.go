package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateRuleResponse struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty" xml:"permissions"`

	// 规则ID
	RuleId *int32 `json:"rule_id,omitempty" xml:"rule_id"`

	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用ID
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 规则状态 0-启用 1-停用
	Status *UpdateRuleResponseStatus `json:"status,omitempty" xml:"status"`

	// 数据解析状态，ENABLE时data_parsing必填 0-启用 1-停用
	DataParsingStatus *UpdateRuleResponseDataParsingStatus `json:"data_parsing_status,omitempty" xml:"data_parsing_status"`

	// SQL查询字段
	SqlField *string `json:"sql_field,omitempty" xml:"sql_field"`

	// SQL查询条件
	SqlWhere *string `json:"sql_where,omitempty" xml:"sql_where"`

	// 完整的规则表达式
	RuleExpress *string `json:"rule_express,omitempty" xml:"rule_express"`

	CreatedUser *CreatedUser `json:"created_user,omitempty" xml:"created_user"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty" xml:"last_updated_user"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty" xml:"created_datetime"`

	// 最后修改时间，timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty" xml:"last_updated_datetime"`
	HttpStatusCode      int    `json:"-"`
}

func (o UpdateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRuleResponse", string(data)}, " ")
}

type UpdateRuleResponseStatus struct {
	value int32
}

type UpdateRuleResponseStatusEnum struct {
	E_0 UpdateRuleResponseStatus
	E_1 UpdateRuleResponseStatus
}

func GetUpdateRuleResponseStatusEnum() UpdateRuleResponseStatusEnum {
	return UpdateRuleResponseStatusEnum{
		E_0: UpdateRuleResponseStatus{
			value: 0,
		}, E_1: UpdateRuleResponseStatus{
			value: 1,
		},
	}
}

func (c UpdateRuleResponseStatus) Value() int32 {
	return c.value
}

func (c UpdateRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateRuleResponseDataParsingStatus struct {
	value int32
}

type UpdateRuleResponseDataParsingStatusEnum struct {
	E_0 UpdateRuleResponseDataParsingStatus
	E_1 UpdateRuleResponseDataParsingStatus
}

func GetUpdateRuleResponseDataParsingStatusEnum() UpdateRuleResponseDataParsingStatusEnum {
	return UpdateRuleResponseDataParsingStatusEnum{
		E_0: UpdateRuleResponseDataParsingStatus{
			value: 0,
		}, E_1: UpdateRuleResponseDataParsingStatus{
			value: 1,
		},
	}
}

func (c UpdateRuleResponseDataParsingStatus) Value() int32 {
	return c.value
}

func (c UpdateRuleResponseDataParsingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleResponseDataParsingStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
