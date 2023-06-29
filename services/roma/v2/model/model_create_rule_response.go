package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRuleResponse Response Object
type CreateRuleResponse struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty"`

	// 规则ID
	RuleId *int32 `json:"rule_id,omitempty"`

	// 规则名称，支持英文大小写，数字，下划线和中划线,长度1-64
	Name *string `json:"name,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 规则状态 0-启用 1-停用
	Status *CreateRuleResponseStatus `json:"status,omitempty"`

	// 数据解析状态，ENABLE时data_parsing必填 0-启用 1-停用
	DataParsingStatus *CreateRuleResponseDataParsingStatus `json:"data_parsing_status,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty"`

	// 最后修改时间，timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o CreateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleResponse", string(data)}, " ")
}

type CreateRuleResponseStatus struct {
	value int32
}

type CreateRuleResponseStatusEnum struct {
	E_0 CreateRuleResponseStatus
	E_1 CreateRuleResponseStatus
}

func GetCreateRuleResponseStatusEnum() CreateRuleResponseStatusEnum {
	return CreateRuleResponseStatusEnum{
		E_0: CreateRuleResponseStatus{
			value: 0,
		}, E_1: CreateRuleResponseStatus{
			value: 1,
		},
	}
}

func (c CreateRuleResponseStatus) Value() int32 {
	return c.value
}

func (c CreateRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateRuleResponseDataParsingStatus struct {
	value int32
}

type CreateRuleResponseDataParsingStatusEnum struct {
	E_0 CreateRuleResponseDataParsingStatus
	E_1 CreateRuleResponseDataParsingStatus
}

func GetCreateRuleResponseDataParsingStatusEnum() CreateRuleResponseDataParsingStatusEnum {
	return CreateRuleResponseDataParsingStatusEnum{
		E_0: CreateRuleResponseDataParsingStatus{
			value: 0,
		}, E_1: CreateRuleResponseDataParsingStatus{
			value: 1,
		},
	}
}

func (c CreateRuleResponseDataParsingStatus) Value() int32 {
	return c.value
}

func (c CreateRuleResponseDataParsingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRuleResponseDataParsingStatus) UnmarshalJSON(b []byte) error {
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
